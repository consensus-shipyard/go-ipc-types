package voting

import (
	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/utils"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"
	"github.com/ipfs/go-cid"
	xerrors "golang.org/x/xerrors"
)

const bitWidth = 5

type Ratio struct {
	Num   uint64
	Denom uint64
}

// Voting is the Go implementation of the Voting<T> struct
// from Rust ipc-actors
type Voting struct {
	GenesisEpoch         abi.ChainEpoch
	SubmissionPeriod     abi.ChainEpoch
	LastVotingExecuted   abi.ChainEpoch
	ExecutableEpochQueue []abi.ChainEpoch //Option<BTreeSet<ChainEpoch>>
	EpochVoteSubmission  cid.Cid          //TCid<THamt<ChainEpoch, EpochVoteSubmissions<T>>>
	Ratio                Ratio
}

func NewWithRatio(store adt.Store, genesis, period abi.ChainEpoch, ratio Ratio) (Voting, error) {
	emptyMapCid, err := adt.StoreEmptyMap(store, bitWidth)
	if err != nil {
		return Voting{}, xerrors.Errorf("failed to create empty map: %w", err)
	}
	return Voting{
		GenesisEpoch:         genesis,
		SubmissionPeriod:     period,
		LastVotingExecuted:   0,
		ExecutableEpochQueue: make([]abi.ChainEpoch, 0),
		EpochVoteSubmission:  emptyMapCid,
		Ratio:                ratio,
	}, nil
}

// EpochVoteSubmissions tracks all the vote submissions of an epoch
// for a checkpoint voting.
type EpochVoteSubmissions struct {
	TotalSubmissionWeight abi.TokenAmount
	MostVotedKey          []byte
	Submitters            cid.Cid // TCid<THamt<Address, ()>>
	SubmissionWeights     cid.Cid // TCid<THamt<UniqueBytesKey, TokenAmount>>
	Submissions           cid.Cid // TCid<THamt<UniqueBytesKey, T>>
}

// ValidatorHasVoted checks if a validator has already voted for a checkpoint.
//
// This function expects the ID (f0 address) of the validator address, which is the
// one used to index validators in the actor state. f1 and f3 addresses will have to
// be translated to f0 addresses before using this function.
func (v *Voting) ValidatorHasVoted(s adt.Store, epoch abi.ChainEpoch, validator address.Address) (bool, error) {
	sub, found, err := utils.GetOutOfHamt[EpochVoteSubmissions](v.EpochVoteSubmission, s, sdk.EpochKey(epoch))
	if err != nil {
		return false, err
	}
	// if no submission found we don't error, we consider that no validator
	// has voted yet.
	if !found {
		return false, nil
	}
	_, found, err = utils.GetOutOfHamt[interface{}](sub.Submitters, s, abi.AddrKey(validator))
	if err != nil {
		return false, err
	}
	return found, nil
}
