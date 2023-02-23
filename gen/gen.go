package main

import (
	gen "github.com/whyrusleeping/cbor-gen"

	"github.com/consensus-shipyard/go-ipc-types/gateway"
	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/subnetactor"
)

func main() {
	// gateway types
	if err := gen.WriteTupleEncodersToFile("./gateway/cbor_gen.go", "gateway",
		gateway.State{},
		gateway.ConstructParams{},
		gateway.Checkpoint{},
		gateway.CheckData{},
		gateway.ChildCheck{},
		gateway.CrossMsgMeta{},
		gateway.StorableMsg{},
		gateway.Subnet{},
		gateway.RawBytes{},
		gateway.CrossMsg{},
		gateway.FundParams{},
		gateway.CrossMsgParams{},
		gateway.ApplyMsgParams{},
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
		subnetactor.Validator{},
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
}
