package ipc_types

//go:generate go run ./gen/gen.go

import (
	"github.com/filecoin-project/go-address"
)

type IPCAddress struct {
	SubnetID   SubnetID
	RawAddress address.Address
}

type SubnetID struct {
	Parent string
	Actor  address.Address
}
