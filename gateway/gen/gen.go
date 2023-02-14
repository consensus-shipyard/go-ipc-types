package main

import (
	"github.com/consensus-shipyard/go-ipc-types/gateway"
	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "gateway",
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
}
