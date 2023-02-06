package main

import (
	"github.com/consensus-shipyard/go-ipc-types/types"
	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "types",
		types.IPCAddress{},
		types.SubnetID{},
		types.IPCSubnetActorState{},
		types.ConstructParams{},
		types.JoinParams{},
		types.Votes{},
		types.Checkpoint{},
		types.CheckData{},
		types.ChildCheck{},
		types.CrossMsgMeta{},
		types.IPCGatewayState{},
		types.ConstructorParams{},
		types.StorableMsg{},
		types.RawBytes{},
		types.CrossMsg{},
		types.FundParams{},
		types.CrossMsgParams{},
		types.ApplyMsgParams{},
		types.CrossMsgs{},
		types.Validator{},
	); err != nil {
		panic(err)
	}
}
