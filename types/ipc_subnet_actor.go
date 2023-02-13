package types

import (
	mbig "math/big"

	"github.com/consensus-shipyard/go-ipc-types/utils"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"
)

type IPCSubnetActorState struct { // TODO change name for IPCSubnetActorState once put in its own file
	Name              string
	ParentID          SubnetID
	IPCGatewayAddr    address.Address
	Consensus         ConsensusType
	MinValidatorStake abi.TokenAmount
	TotalStake        abi.TokenAmount
	Stake             cid.Cid // TCid<THamt<Cid,TokenAmount>>
	Status            Status
	Genesis           []byte
	FinalityThreshold ChainEpoch
	CheckPeriod       ChainEpoch
	Checkpoints       cid.Cid // TCid<THamt<Cid, Checkpoint>>
	WindowChecks      cid.Cid // TCid<THamt<Cid, Votes>>,
	ValidatorSet      []Validator
	MinValidators     uint64
}

func (st *IPCSubnetActorState) GetStake(s adt.Store, id address.Address) (*abi.TokenAmount, error) {
	tokenAmount, err := utils.GetOutOfHamt[abi.TokenAmount](st.Stake, s, abi.AddrKey(id))
	return tokenAmount, err
}

func (st *IPCSubnetActorState) GetCheckpoint(s adt.Store, id address.Address) (*Checkpoint, error) {
	checkpoint, err := utils.GetOutOfHamt[Checkpoint](st.Stake, s, abi.AddrKey(id))
	return checkpoint, err
}

func (st *IPCSubnetActorState) GetWindowCheck(s adt.Store, id address.Address) (*Votes, error) {
	votes, err := utils.GetOutOfHamt[Votes](st.Stake, s, abi.AddrKey(id))
	return votes, err
}

func (st *IPCSubnetActorState) GetCrossMsgs(s adt.Store, c cid.Cid) (*CrossMsgs, error) {
	crossMsgs, err := utils.GetOutOfHamt[CrossMsgs](st.Stake, s, abi.CidKey(c))
	return crossMsgs, err
}

func (st *IPCSubnetActorState) HasMajorityVote(s adt.Store, v Votes) (bool, error) {
	sum := big.Zero()
	for _, m := range v.Validators {
		stake, err := st.GetStake(s, m)
		if err != nil {
			return false, err
		}
		sum = big.Sum(sum, *stake)
	}
	fsum := new(mbig.Rat).SetInt(sum.Int)
	fTotal := new(mbig.Rat).SetInt(st.TotalStake.Int)
	div := new(mbig.Rat).SetFrac(fsum.Num(), fTotal.Num())
	threshold := utils.MajorityThreshold()
	return div.Cmp(&threshold) >= 0, nil

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
	MinValidatorStake abi.TokenAmount
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

func (c *Checkpoint) CrossMsgMeta(from, to *SubnetID) (*CrossMsgMeta, bool) {
	for i, m := range c.Data.CrossMsgs {
		if *from == m.From && *to == m.To {
			return &c.Data.CrossMsgs[i], true
		}
	}
	return nil, false
}

func (s *Checkpoint) CrossMsgMetaIndex(from, to *SubnetID) (int, bool) {
	for i, m := range s.Data.CrossMsgs {
		if *from == m.From && *to == m.To {
			return i, true
		}
	}
	return 0, false
}

func CheckpointEpoch(epoch, period ChainEpoch) ChainEpoch {
	return (epoch / period) * period
}

func WindowEpoch(epoch, period ChainEpoch) ChainEpoch {
	ind := epoch / period
	return period * (ind + 1)
}

type CheckData struct {
	Source    SubnetID
	TipSet    []byte
	Epoch     ChainEpoch
	PrevCheck cid.Cid // TCid<TLink<Checkpoint>>
	Children  []ChildCheck
	CrossMsgs []CrossMsgMeta
}

type ChildCheck struct {
	Source SubnetID
	Checks cid.Cid // Vec<TCid<TLink<Checkpoint>>>,
}

type CrossMsgMeta struct {
	From    SubnetID
	To      SubnetID
	MsgsCID cid.Cid // TCid<TLink<CrossMsgs>>,
	Nonce   uint64
	Value   abi.TokenAmount
}
