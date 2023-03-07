package subnetactor

import (
	mbig "math/big"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"

	"github.com/consensus-shipyard/go-ipc-types/gateway"
	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/utils"
	"github.com/consensus-shipyard/go-ipc-types/validator"
)

type State struct {
	Name              string
	ParentID          sdk.SubnetID
	IPCGatewayAddr    address.Address
	Consensus         ConsensusType
	MinValidatorStake abi.TokenAmount
	TotalStake        abi.TokenAmount
	Stake             cid.Cid // TCid<THamt<Cid,TokenAmount>>
	Status            sdk.Status
	Genesis           []byte
	FinalityThreshold abi.ChainEpoch
	CheckPeriod       abi.ChainEpoch
	Checkpoints       cid.Cid // TCid<THamt<Cid, Checkpoint>>
	WindowChecks      cid.Cid // TCid<THamt<Cid, Votes>>,
	ValidatorSet      *validator.Set
	MinValidators     uint64
}

func (st *State) GetStake(s adt.Store, id address.Address) (abi.TokenAmount, error) {
	out, found, err := utils.GetOutOfHamt[abi.TokenAmount](st.Stake, s, abi.AddrKey(id))
	if err != nil {
		return abi.NewTokenAmount(0), err
	}
	if !found {
		return abi.NewTokenAmount(0), err
	}
	return *out, err
}

func (st *State) GetCheckpoint(s adt.Store, id address.Address) (*gateway.Checkpoint, bool, error) {
	return utils.GetOutOfHamt[gateway.Checkpoint](st.Stake, s, abi.AddrKey(id))
}

func (st *State) HasMajorityVote(s adt.Store, v Votes) (bool, error) {
	sum := big.Zero()
	for _, m := range v.Validators {
		stake, err := st.GetStake(s, m)
		if err != nil {
			return false, err
		}
		sum = big.Sum(sum, stake)
	}
	fsum := new(mbig.Rat).SetInt(sum.Int)
	fTotal := new(mbig.Rat).SetInt(st.TotalStake.Int)
	div := new(mbig.Rat).SetFrac(fsum.Num(), fTotal.Num())
	threshold := utils.MajorityThreshold()
	return div.Cmp(&threshold) >= 0, nil
}
