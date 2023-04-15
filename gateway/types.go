package gateway

import (
	"bytes"
	"fmt"

	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	xerrors "golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"

	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/utils"
)

// ManifestID is the id used to index the gateway actor
// in the builtin-actors bundle.
const ManifestID = "ipc_gateway"

type Subnet struct {
	ID                   sdk.SubnetID
	Stake                abi.TokenAmount
	TopDownMsgs          cid.Cid // TCid<TAmt<CrossMsg, CROSSMSG_AMT_BITWIDTH>>,
	Nonce                uint64
	CircSupply           abi.TokenAmount
	Status               sdk.Status
	PrevCheckpoint       *BottomUpCheckpoint
	AppliedBottomupNonce uint64
}

func (sn *Subnet) GetTopDownMsg(s adt.Store, nonce uint64) (*CrossMsg, bool, error) {
	return utils.GetOutOfArray[CrossMsg](sn.TopDownMsgs, s, nonce, CrossMsgsAMTBitwidth)
}

// TopDownMsgFromNonce gets the latest topDownMessages from a specific nonce
// (including the one specified, i.e. [nonce, latest], both limits
// included).
func (sn *Subnet) TopDownMsgsFromNonce(s adt.Store, nonce uint64) ([]*CrossMsg, error) {
	crossMsgs, err := adt.AsArray(s, sn.TopDownMsgs, CrossMsgsAMTBitwidth)
	if err != nil {
		return nil, xerrors.Errorf("failed to load cross-msgs: %w", err)
	}
	// TODO: Consider setting the length of the slice in advance
	// to improve performance.
	out := make([]*CrossMsg, 0)
	for i := nonce; i < sn.Nonce; i++ {
		msg, found, err := getTopDownMsg(crossMsgs, i)
		if err != nil {
			return nil, err
		}
		if found {
			out = append(out, msg)
		}
	}
	return out, nil
}

func getTopDownMsg(crossMsgs *adt.Array, nonce uint64) (*CrossMsg, bool, error) {
	var out CrossMsg
	found, err := crossMsgs.Get(nonce, &out)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get cross-msg with nonce %v: %w", nonce, err)
	}
	if !found {
		return nil, found, nil
	}
	return &out, found, nil
}

type StorableMsg struct {
	From   sdk.IPCAddress
	To     sdk.IPCAddress
	Method abi.MethodNum
	Params []byte
	Value  abi.TokenAmount
	Nonce  uint64
}

func (sm *StorableMsg) IPCType() IPCMsgType {
	toSubnetID := sm.To.SubnetID
	fromSubnetID := sm.From.SubnetID
	if sdk.IsBottomUp(fromSubnetID, toSubnetID) {
		return IPCMsgTypeBottomUp
	}
	return IPCMsgTypeTopDown
}

const (
	IPCMsgTypeUnknown IPCMsgType = iota
	IPCMsgTypeBottomUp
	IPCMsgTypeTopDown
)

type IPCMsgType int

type ConstructParams struct {
	NetworkName      string
	CheckpointPeriod abi.ChainEpoch
}

type CrossMsg struct {
	Msg     StorableMsg
	Wrapped bool
}

type FundParams struct {
	Value abi.TokenAmount
}

type CrossMsgParams struct {
	CrossMsg    CrossMsg
	Destination sdk.SubnetID
}

type CrossMsgs struct {
	Msgs []CrossMsg
}

const CrossMsgsAMTBitwidth = 3

type BottomUpCheckpoint struct {
	Data CheckData
	Sig  []byte
}

func NewBottomUpCheckpoint(subnet sdk.SubnetID, epoch abi.ChainEpoch) *BottomUpCheckpoint {
	return &BottomUpCheckpoint{Data: CheckData{Source: subnet, Epoch: epoch}}
}

func (c *BottomUpCheckpoint) Cid() (cid.Cid, error) {
	buf := new(bytes.Buffer)
	if err := c.Data.MarshalCBOR(buf); err != nil {
		return cid.Undef, err
	}
	h, err := mh.Sum(buf.Bytes(), abi.HashFunction, -1)
	if err != nil {
		return cid.Undef, err
	}

	return cid.NewCidV1(abi.CidBuilder.GetCodec(), h), nil
}

func CheckpointEpoch(epoch, period abi.ChainEpoch) abi.ChainEpoch {
	return (epoch / period) * period
}

func WindowEpoch(epoch, period abi.ChainEpoch) abi.ChainEpoch {
	ind := epoch / period
	return period * (ind + 1)
}

type CheckData struct {
	Source    sdk.SubnetID
	Proof     []byte
	Epoch     abi.ChainEpoch
	PrevCheck cid.Cid // TCid<TLink<Checkpoint>>
	Children  []ChildCheck
	CrossMsgs BatchCrossMsgs
}

type ChildCheck struct {
	Source sdk.SubnetID
	Checks []cid.Cid // Vec<TCid<TLink<Checkpoint>>>,
}

type BatchCrossMsgs struct {
	CrossMsgs []CrossMsg
	Fee       abi.TokenAmount
}

type TopDownCheckpoint struct {
	Epoch       abi.ChainEpoch
	TopDownMsgs []CrossMsgs
}
