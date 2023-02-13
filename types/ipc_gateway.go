package types

import (
	"fmt"
	"strings"

	"github.com/consensus-shipyard/go-ipc-types/utils"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"
	"github.com/ipfs/go-cid"
)

type IPCGatewayState struct {
	NetworkName          SubnetID
	TotalSubnets         uint64
	MinStake             abi.TokenAmount
	Subnets              cid.Cid // TCid<THamt<Cid, Subnet>>
	CheckPeriod          ChainEpoch
	Checkpoints          cid.Cid //TCid<THamt<ChainEpoch, Checkpoint>>
	CheckMsgRegistry     cid.Cid //TCid<THamt<TCid<TLink<CrossMsgs>>, CrossMsgs>>
	Postbox              cid.Cid // TCid<THamt<Cid, Vec<u8>>>;
	Nonce                uint64
	BottomupNonce        uint64
	BottomupMsgMeta      cid.Cid // TCid<TAmt<CrossMsgMeta, CROSSMSG_AMT_BITWIDTH>>
	AppliedBottomupNonce uint64
	AppliedTopdownNonce  uint64
}

type Subnet struct {
	ID             SubnetID
	Stake          abi.TokenAmount
	TopDownMsgs    cid.Cid // TCid<TAmt<CrossMsg, CROSSMSG_AMT_BITWIDTH>>,
	Nonce          uint64
	CircSupply     abi.TokenAmount
	Status         Status
	PrevCheckpoint Checkpoint
}

func (sn Subnet) GetTopDownMsg(s adt.Store, nonce uint64) (*CrossMsg, error) {
	return utils.GetOutOfArray[CrossMsg](sn.TopDownMsgs, s, nonce, CrossMsgsAMTBitwidth)
}

func GetTopDownMsg(crossMsgs *adt.Array, nonce uint64) (*CrossMsg, error) {
	var out CrossMsg
	found, err := crossMsgs.Get(nonce, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get cross-msg with nonce %v: %w", nonce, err)
	}
	if !found {
		return nil, nil
	}
	return &out, nil
}

func (st *IPCGatewayState) GetSubnet(s adt.Store, id SubnetID) (*Subnet, error) {
	key, err := abi.ParseUIntKey(id.String())
	id.Bytes()
	if err != nil {
		return nil, err
	}
	subnet, err := utils.GetOutOfHamt[Subnet](st.Subnets, s, abi.UIntKey(key))
	return subnet, err
}

func (st *IPCGatewayState) GetCheckpoints(s adt.Store, c ChainEpoch) (*Checkpoint, error) {
	checkpoint, err := utils.GetOutOfHamt[Checkpoint](st.Checkpoints, s, abi.UIntKey(uint64(c)))
	return checkpoint, err
}

func (st *IPCGatewayState) GetCrossMsgs(s adt.Store, cID cid.Cid) (*CrossMsgs, error) {
	crossMsgs, err := utils.GetOutOfHamt[CrossMsgs](st.Checkpoints, s, abi.CidKey(cID))
	return crossMsgs, err
}

func (st *IPCGatewayState) GetBottomUpMsgMeta(s adt.Store, cID cid.Cid, nonce uint64) (*CrossMsgMeta, error) {
	return utils.GetOutOfArray[CrossMsgMeta](cID, s, nonce, CrossMsgsAMTBitwidth)
}

func (st *IPCGatewayState) GetTopDownMsg(s adt.Store, id SubnetID, nonce uint64) (*CrossMsg, error) {
	sh, err := st.GetSubnet(s, id)
	if err != nil {
		return nil, err
	}
	CrossMsg, err := sh.GetTopDownMsg(s, nonce)
	return CrossMsg, err
}

func IsBottomup(from SubnetID, to SubnetID) bool {
	ptrSubnetID, index := from.CommonParent(to)
	if ptrSubnetID == nil {
		return false
	}

	a := from.String()
	components := strings.Split(a, SubnetSeparator)
	count := len(components) - 1
	return count > index

}

// BottomUpMsgFromNonce gets the latest bottomUpMetas from a specific nonce
// (including the one specified, i.e. [nonce, latest], both limits
// included).
func (st *IPCGatewayState) BottomUpMsgFromNonce(s adt.Store, nonce uint64) ([]*CrossMsgMeta, error) {
	out := make([]*CrossMsgMeta, 0)
	adtArray, err := adt.AsArray(s, st.BottomupMsgMeta, CrossMsgsAMTBitwidth)
	if err != nil {
		return nil, err
	}
	for i := nonce; i < st.BottomupNonce; i++ {
		meta, err := utils.GetOutOfAdtArray[CrossMsgMeta](adtArray, i)
		if err != nil {
			return nil, err
		}
		if meta != nil { // then found
			out = append(out, meta)
		}
	}
	return out, nil
}

type StorableMsg struct {
	From   IPCAddress
	To     IPCAddress
	Method MethodNum
	Params RawBytes
	Value  abi.TokenAmount
	Nonce  uint64
}

func (sm *StorableMsg) IPCType() IPCMsgType {
	toSubnetID := sm.To.SubnetID
	fromSubnetID := sm.From.SubnetID
	if IsBottomup(fromSubnetID, toSubnetID) {
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

type ConstructorParams struct {
	NetworkName      string
	CheckpointPeriod ChainEpoch
}

type MethodNum uint64

type RawBytes struct {
	Bytes []byte
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
	Destination SubnetID
}

type ApplyMsgParams struct {
	CrossMsg CrossMsg
}

type CrossMsgs struct {
	Msgs  []CrossMsg
	Metas []CrossMsgMeta
}

const CrossMsgsAMTBitwidth = 3
