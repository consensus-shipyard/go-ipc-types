package main

import (
	ipcsdk "github.com/consensus-shipyard/go-ipc-types/ipcsdk"
	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "ipcsdk",
		ipcsdk.IPCAddress{},
		ipcsdk.SubnetID{},
	); err != nil {
		panic(err)
	}
}
