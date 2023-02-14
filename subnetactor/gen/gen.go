package main

import (
	subnetactor "github.com/consensus-shipyard/go-ipc-types/subnetactor"
	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "subnetactor",
		subnetactor.State{},
		subnetactor.ConstructParams{},
		subnetactor.JoinParams{},
		subnetactor.Votes{},
		subnetactor.Validator{},
	); err != nil {
		panic(err)
	}
}
