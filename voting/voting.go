package voting

import (
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
