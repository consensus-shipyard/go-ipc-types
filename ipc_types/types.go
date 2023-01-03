package ipc_types

//go:generate go run ./gen/gen.go

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"math/big"
)

// Go types for ipc-gateway/src/subnetid.rs, ipc-gateway/src/address.rs and ipc-subnet-actor and ipc-gateways' states and params

// TODO ----------------------------------------------------------------------
// TODO here below subnetid.rs and address.rs and related
// TODO if moving this code to IPC gateway this could be taken from the same constant as written in Rust

type IPCAddress struct {
	SubnetID   SubnetID
	RawAddress address.Address
}

type SubnetID struct {
	Parent string
	Actor  address.Address
}

// ----------------------------------------------------
// TODO Move them in separate files that somewhat follow the division in Rust
// TODO here below ipc-subnet-actor/src/state.rs and ipc-subnet-actor/src/types.rs (separate further?), and other types associated to them
// ------------------------------------------------------
type IPCSubnetActorState struct { // TODO change name for State once put in its own file
	Name              string
	ParentID          SubnetID
	IpcGatewayAddr    address.Address
	Consensus         ConsensusType
	MinValidatorStake TokenAmount
	TotalStake        TokenAmount
	Stake             cid.Cid //TCid<THamt<Cid,TokenAmount>>
	Status            Status
	Genesis           []byte
	FinalityThreshold ChainEpoch
	CheckPeriod       ChainEpoch
	Checkpoints       cid.Cid //TCid<THamt<Cid, Checkpoint>>
	WindowChecks      cid.Cid //TCid<THamt<Cid, Votes>>,
	ValidatorSet      []Validator
	MinValidators     uint64
}

type ConsensusType int64

const (
	Delegated ConsensusType = iota
	PoW
	Tendermint
	Mir
	FilecoinEC
	Dummy
)

type TokenAmount struct {
	Atto big.Int
}

type Status int64

const (
	Instantiated Status = iota
	Active
	Inactive
	Terminating
	Killed
)

type ChainEpoch int64

type Validator struct {
	Addr    address.Address
	NetAddr string
}

type ConstructParams struct {
	Parent            SubnetID
	Name              string
	IpcGatewayAddr    uint64
	Consensus         ConsensusType
	MinValidatorStake TokenAmount
	MinValidators     uint64
	FinalityThreshold ChainEpoch
	CheckPeriod       ChainEpoch
	Genesis           []byte
}

type JoinParams struct {
	ValidatorNetAddr string
}

type Votes struct {
	Validators []address.Address
}

type Checkpoint struct {
	Data CheckData
	Sig  []byte
}

type CheckData struct {
	Source    SubnetID
	TipSet    []byte
	Epoch     ChainEpoch
	PrevCheck cid.Cid //TCid<TLink<Checkpoint>>
	Children  []ChildCheck
	CrossMsgs []CrossMsgMeta
}

type ChildCheck struct {
	Source SubnetID
	Checks cid.Cid //Vec<TCid<TLink<Checkpoint>>>,
}

type CrossMsgMeta struct {
	From    SubnetID
	To      SubnetID
	MsgsCID cid.Cid //TCid<TLink<CrossMsgs>>,
	Nonce   uint64
	Value   TokenAmount
}

// ----------------------------------------------------
// TODO Move them in separate files that somewhat follow the division in Rust
// TODO here below  ipc-gateways state and params (and other types associated to them)(separate further?)
// ------------------------------------------------------
type IPCGatewayState struct { //TODO change name to state once in its own file
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
