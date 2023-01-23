package main

import (
	gen "github.com/whyrusleeping/cbor-gen"
	"go-ipc-types/types"
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
