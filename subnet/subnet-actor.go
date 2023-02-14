package subnetactor

//go:generate go run ./gen/gen.go

import (
	mbig "math/big"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"

	"github.com/consensus-shipyard/go-ipc-types/gateway"
	"github.com/consensus-shipyard/go-ipc-types/ipcsdk"
	"github.com/consensus-shipyard/go-ipc-types/utils"
)

type State struct {
	Name              string
	ParentID          ipcsdk.SubnetID
	IPCGatewayAddr    address.Address
	Consensus         ConsensusType
	MinValidatorStake abi.TokenAmount
	TotalStake        abi.TokenAmount
	Stake             cid.Cid // TCid<THamt<Cid,TokenAmount>>
	Status            ipcsdk.Status
	Genesis           []byte
	FinalityThreshold abi.ChainEpoch
	CheckPeriod       abi.ChainEpoch
	Checkpoints       cid.Cid // TCid<THamt<Cid, Checkpoint>>
	WindowChecks      cid.Cid // TCid<THamt<Cid, Votes>>,
	ValidatorSet      []Validator
	MinValidators     uint64
}

func (st *State) GetStake(s adt.Store, id address.Address) (*abi.TokenAmount, error) {
	tokenAmount, err := utils.GetOutOfHamt[abi.TokenAmount](st.Stake, s, abi.AddrKey(id))
	return tokenAmount, err
}

func (st *State) GetCheckpoint(s adt.Store, id address.Address) (*gateway.Checkpoint, error) {
	checkpoint, err := utils.GetOutOfHamt[gateway.Checkpoint](st.Stake, s, abi.AddrKey(id))
	return checkpoint, err
}

func (st *State) GetWindowCheck(s adt.Store, id address.Address) (*Votes, error) {
	votes, err := utils.GetOutOfHamt[Votes](st.Stake, s, abi.AddrKey(id))
	return votes, err
}

func (st *State) GetCrossMsgs(s adt.Store, c cid.Cid) (*gateway.CrossMsgs, error) {
	crossMsgs, err := utils.GetOutOfHamt[gateway.CrossMsgs](st.Stake, s, abi.CidKey(c))
	return crossMsgs, err
}

func (st *State) HasMajorityVote(s adt.Store, v Votes) (bool, error) {
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
