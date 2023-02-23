package main

import (
	gen "github.com/whyrusleeping/cbor-gen"

	"github.com/consensus-shipyard/go-ipc-types/sdk"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "sdk",
		sdk.IPCAddress{},
		sdk.SubnetID{},
	); err != nil {
		panic(err)
	}
}
