package ipc_types

import "github.com/ipfs/go-cid"

type IPCGatewayState struct {
	NetworkName          SubnetID
	TotalSubnets         uint64
	MinStake             TokenAmount
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

type ConstructorParams struct {
	NetworkName      string
	CheckpointPeriod ChainEpoch
}

type StorableMsg struct {
	From   IPCAddress
	To     IPCAddress
	Method MethodNum
	Params RawBytes
	Value  TokenAmount
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
	Value TokenAmount
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
