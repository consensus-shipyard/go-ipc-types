package main

import (
	gen "github.com/whyrusleeping/cbor-gen"

	"github.com/consensus-shipyard/go-ipc-types/gateway"
	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/subnetactor"
	"github.com/consensus-shipyard/go-ipc-types/validator"
	"github.com/consensus-shipyard/go-ipc-types/voting"
)

func main() {
	// gateway types
	if err := gen.WriteTupleEncodersToFile("./gateway/cbor_gen.go", "gateway",
		gateway.State{},
		gateway.ConstructParams{},
		gateway.BottomUpCheckpoint{},
		gateway.TopDownCheckpoint{},
		gateway.CheckData{},
		gateway.ChildCheck{},
		gateway.BatchCrossMsgs{},
		gateway.StorableMsg{},
		gateway.Subnet{},
		gateway.CrossMsg{},
		gateway.FundParams{},
		gateway.ReleaseParams{},
		gateway.AmountParams{},
		gateway.InitGenesisEpochParams{},
		gateway.CrossMsgParams{},
		gateway.CrossMsgs{},
	); err != nil {
		panic(err)
	}

	// subnet actor types
	if err := gen.WriteTupleEncodersToFile("./subnetactor/cbor_gen.go", "subnetactor",
		subnetactor.State{},
		subnetactor.ConstructParams{},
		subnetactor.JoinParams{},
		subnetactor.Votes{},
	); err != nil {
		panic(err)
	}

	// sdk types
	if err := gen.WriteTupleEncodersToFile("./sdk/cbor_gen.go", "sdk",
		sdk.IPCAddress{},
		sdk.SubnetID{},
	); err != nil {
		panic(err)
	}

	// validator types
	if err := gen.WriteTupleEncodersToFile("./validator/cbor_gen.go", "validator",
		validator.Validator{},
		validator.Set{},
		validator.OnChainValidators{},
	); err != nil {
		panic(err)
	}

	// common types
	if err := gen.WriteTupleEncodersToFile("./voting/cbor_gen.go", "voting",
		voting.Voting{},
		voting.Ratio{},
		voting.EpochVoteSubmissions{},
	); err != nil {
		panic(err)
	}
}
