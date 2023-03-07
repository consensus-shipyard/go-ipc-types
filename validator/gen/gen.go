package main

import (
	gen "github.com/whyrusleeping/cbor-gen"

	"github.com/consensus-shipyard/go-ipc-types/validator"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "validator",
		validator.Validator{},
		validator.Set{},
	); err != nil {
		panic(err)
	}
}
