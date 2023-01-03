package ipc_types

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"math/big"
)

type IPCSubnetActorState struct { // TODO change name for IPCSubnetActorState once put in its own file
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
