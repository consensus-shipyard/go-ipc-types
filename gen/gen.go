package main

import (
	"github.com/consensus-shipyard/go-ipc-types/gateway"
	"github.com/consensus-shipyard/go-ipc-types/ipcsdk"
	"github.com/consensus-shipyard/go-ipc-types/subnetactor"
	gen "github.com/whyrusleeping/cbor-gen"
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
	if err := gen.WriteTupleEncodersToFile("./ipcsdk/cbor_gen.go", "ipcsdk",
		ipcsdk.IPCAddress{},
		ipcsdk.SubnetID{},
	); err != nil {
		panic(err)
	}
}