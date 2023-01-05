package types

//go:generate go run ./gen/gen.go

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"path"
	"path/filepath"
	"strings"
)

type IPCAddress struct {
	SubnetID   SubnetID
	RawAddress address.Address
}

func IPCAddressFromString(addr string) (*IPCAddress, error) {
	r := strings.Split(addr, IPCAddrSeparator)
	if len(r) != 2 {
		return nil, fmt.Errorf("invalid IPCAddress string type") // TODO Create new error, TODO Define Undef for IPCAddress
	}
	rawAddress, err := address.NewFromString(r[1])
	if err != nil {
		return nil, err
	}
	ptrSubnetID, err := SubnetIDFromString(r[0])
	if err != nil {
		return nil, err
	}
	return &IPCAddress{*ptrSubnetID, rawAddress}, nil
}

type SubnetID struct {
	Parent string
	Actor  address.Address
}

var id0, _ = address.NewIDAddress(0) //TODO check with Alfonso

const (
	RootStr          = "/root"
	SubnetSeparator  = "/"
	UndefStr         = ""
	IPCAddrSeparator = ":"
)

// RootSubnet is the ID of the root network
var RootSubnet = SubnetID{
	Parent: RootStr,
	Actor:  id0,
}

// UndefSubnetID is the undef ID
var UndefSubnetID = SubnetID{
	Parent: SubnetSeparator,
	Actor:  id0,
}

func NewSubnetIDFromString(addr string) (*SubnetID, error) {
	var out SubnetID
	if addr == RootSubnet.String() {
		out = RootSubnet
		return &out, nil
	}
	dir, file := filepath.Split(addr)
	act, err := address.NewFromString(file)
	if err != nil {
		return nil, err
	}
	return &SubnetID{
		Parent: dir,
		Actor:  act,
	}, nil
}

func NewSubnetID(parent SubnetID, subnetAct address.Address) *SubnetID {
	return &SubnetID{
		parent.String(),
		subnetAct,
	}
}

func (s SubnetID) Bytes() []byte {
	strID := s.String()
	return []byte(strID)
}

// String returns the id in string form.
func (id SubnetID) String() string {
	if id == RootSubnet {
		if id.Parent != UndefStr {
			return RootStr
		} else {
			return UndefStr
		}
	}
	return strings.Join([]string{id.Parent, id.Actor.String()}, SubnetSeparator)
}

func (id SubnetID) CommonParent(other SubnetID) (*SubnetID, int) {
	s1 := strings.Split(id.String(), SubnetSeparator)
	s2 := strings.Split(other.String(), SubnetSeparator)
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	out := SubnetSeparator
	l := 0
	for i, s := range s2 {
		if s == s1[i] {
			out = path.Join(out, s)
			l = i
		} else {
			sn, err := SubnetIDFromString(out)
			if err != nil {
				return nil, 0
			}
			return sn, l
		}
	}
	sn, err := SubnetIDFromString(out)
	if err != nil {
		return nil, 0
	}
	return sn, l
}

func SubnetIDFromString(str string) (*SubnetID, error) {
	switch str {
	case RootStr:
		out := RootSubnet
		return &out, nil
	case UndefStr:
		return nil, nil
	}

	s1 := strings.Split(str, SubnetSeparator)
	actor, err := address.NewFromString(s1[len(s1)-1])
	if err != nil {
		return nil, err
	}
	return &SubnetID{
		Parent: strings.Join(s1[:len(s1)-1], SubnetSeparator),
		Actor:  actor,
	}, nil
}

func (id SubnetID) Down(curr SubnetID) *SubnetID {
	s1 := strings.Split(id.String(), SubnetSeparator)
	s2 := strings.Split(curr.String(), SubnetSeparator)
	// curr needs to be contained in id
	if len(s2) >= len(s1) {
		return nil
	}
	_, l := id.CommonParent(curr)
	out := SubnetSeparator
	for i := 0; i <= l+1 && i < len(s1); i++ {
		if i < len(s2) && s1[i] != s2[i] {
			// they are not in a common path
			return nil
		}
		out = path.Join(out, s1[i])
	}
	sn, err := SubnetIDFromString(out)
	if err != nil {
		return nil
	}
	return sn
}

func (id SubnetID) Up(curr SubnetID) *SubnetID {
	s1 := strings.Split(id.String(), SubnetSeparator)
	s2 := strings.Split(curr.String(), SubnetSeparator)
	// curr needs to be contained in id
	if len(s2) > len(s1) {
		return nil
	}

	_, l := id.CommonParent(curr)
	out := SubnetSeparator
	for i := 0; i < l; i++ {
		if i < len(s1) && s1[i] != s2[i] {
			// they are not in a common path
			return nil
		}
		out = path.Join(out, s1[i])
	}
	sn, err := SubnetIDFromString(out)
	if err != nil {
		return nil
	}
	return sn
}
