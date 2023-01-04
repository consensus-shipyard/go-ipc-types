package ipc_types

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/filecoin-project/specs-actors/v7/actors/builtin"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

type IPCGatewayState struct {
	NetworkName          SubnetID
	TotalSubnets         uint64
	MinStake             abi.TokenAmount
	Subnets              cid.Cid //TCid<THamt<Cid, Subnet>>
	CheckPeriod          ChainEpoch
	Checkpoints          cid.Cid //TCid<THamt<ChainEpoch, Checkpoint>>
	CheckMsgRegistry     cid.Cid //TCid<THamt<TCid<TLink<CrossMsgs>>, CrossMsgs>>
	Nonce                uint64
	BottomupNonce        uint64
	BottomupMsgMeta      cid.Cid //TCid<TAmt<CrossMsgMeta, CROSSMSG_AMT_BITWIDTH>>
	AppliedBottomupNonce uint64
	AppliedTopdownNonce  uint64
}

func (st *IPCGatewayState) GetSubnet(s adt.Store, id address.Address) (SubnetID, error) {
	subnetID, err := getOutOfHamt[SubnetID](st.Subnets, s, abi.AddrKey(id))
	return subnetID, err
}

func (st *IPCGatewayState) GetCheckpoints(s adt.Store, c ChainEpoch) (Checkpoint, error) {
	checkpoint, err := getOutOfHamt[Checkpoint](st.Checkpoints, s, abi.UIntKey(uint64(c)))
	return checkpoint, err
}

func (st *IPCGatewayState) GetCheckMsgRegistry(s adt.Store, cID cid.Cid) (CrossMsgs, error) {
	crossMsgs, err := getOutOfHamt[CrossMsgs](st.Checkpoints, s, abi.CidKey(cID))
	return crossMsgs, err
}

const CrossMsgsAMTBitwidth = 3

func (st *IPCGatewayState) GetBottomUpMsgMeta(s adt.Store, cID cid.Cid, index uint64) (CrossMsgMeta, error) {
	var out CrossMsgMeta
	array, err := adt.AsArray(s, cID, CrossMsgsAMTBitwidth)
	if err != nil {
		return out, err
	}
	_, err = array.Get(index, &out)
	return out, err
}

// getOutOfHamt takes a generic type that must implement cbor.Unmarshaler
// and returns a particular vale of a THamt type passed as cid.Cid given the key
// If the type does not implement cbor.Unmarshaler then this returns a runtime error
func getOutOfHamt[T any](cID cid.Cid, s adt.Store, k abi.Keyer) (T, error) {
	var out T
	adtMap, err := adt.AsMap(s, cID, builtin.DefaultHamtBitwidth)
	if err != nil {
		return out, xerrors.Errorf("failed to get stake: %w", err)
	}
	if i, ok := (any(&out)).(cbor.Unmarshaler); ok {
		_, err = adtMap.Get(k, i)
	} else {
		return out, fmt.Errorf("the type *%T does not implement the cbor.Unmarshaler interface", out)
	}
	return out, err
}

type ConstructorParams struct {
	NetworkName      string
	CheckpointPeriod ChainEpoch
}

type StorableMsg struct {
	From   IPCAddress
	To     IPCAddress
	Method MethodNum
	Params RawBytes
	Value  abi.TokenAmount
	Nonce  uint64
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
