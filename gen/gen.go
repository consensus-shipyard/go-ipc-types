package main

import (
	gen "github.com/whyrusleeping/cbor-gen"
	"go-ipc-types/ipc_types"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "ipc_types",
		ipc_types.IPCAddress{},
		ipc_types.SubnetID{},
		ipc_types.TokenAmount{},
		ipc_types.IPCSubnetActorState{},
		ipc_types.ConstructParams{},
		ipc_types.JoinParams{},
		ipc_types.Votes{},
		ipc_types.Checkpoint{},
		ipc_types.CheckData{},
		ipc_types.ChildCheck{},
		ipc_types.CrossMsgMeta{},
		ipc_types.IPCGatewayState{},
		ipc_types.ConstructorParams{},
		ipc_types.StorableMsg{},
		ipc_types.RawBytes{},
		ipc_types.CrossMsg{},
		ipc_types.FundParams{},
		ipc_types.CrossMsgParams{},
		ipc_types.ApplyMsgParams{},
		ipc_types.CrossMsgs{},
	); err != nil {
		panic(err)
	}
}
