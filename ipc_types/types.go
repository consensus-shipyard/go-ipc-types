package ipc_types

//go:generate go run ./gen/gen.go

import (
	"github.com/filecoin-project/go-address"
	"strings"
)

type IPCAddress struct {
	SubnetID   SubnetID
	RawAddress address.Address
}

type SubnetID struct {
	Parent string
	Actor  address.Address
}

var id0, _ = address.NewIDAddress(0) //TODO check with Alfonso

const (
	RootStr         = "/root"
	SubnetSeparator = "/"
	UndefStr        = ""
	HCAddrSeparator = ":"
)

// RootSubnet is the ID of the root network
var RootSubnet = SubnetID{
	Parent: RootStr,
	Actor:  id0,
}

// DefaultSubnet is the undef ID
var DefaultSubnet = SubnetID{
	Parent: UndefStr,
	Actor:  id0,
}

func NewSubnetIDFromString(parent string, subnetAct address.Address) SubnetID {
	return SubnetID{
		Parent: parent,
		Actor:  subnetAct,
	}
}

func NewSubnetID(parent SubnetID, subnetAct address.Address) SubnetID {
	return NewSubnetIDFromString(parent.String(), subnetAct)
}

func (s *SubnetID) ToBytes() []byte {
	strID := s.String()
	return []byte(strID)
}

// String returns the id in string form.
func (id SubnetID) String() string {
	if id == RootSubnet {
		if id.Parent != UndefStr {
			return RootStr
		} else {
			return UndefStr //TODO found bug in here perhaps? check with Alfonso
		}
	}
	return strings.Join([]string{id.Parent, id.Actor.String()}, SubnetSeparator)
}
