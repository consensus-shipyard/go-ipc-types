package main

import (
	sdk "github.com/consensus-shipyard/go-ipc-types/sdk"
	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "sdk",
		sdk.Address{},
		sdk.SubnetID{},
	); err != nil {
		panic(err)
	}
}
