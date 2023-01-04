package ipc_types

//go:generate go run ./gen/gen.go

import (
	"github.com/filecoin-project/go-address"
	"golang.org/x/xerrors"
	"path"
	"strings"
)

var Undef = IPCAddress{}

type IPCAddress struct {
	SubnetID   SubnetID
	RawAddress address.Address
}

func newIPCAddressFromID(id uint64) IPCAddress {
	d := UndefSubnetID
	rawaddr, _ := address.NewIDAddress(id)
	return IPCAddress{
		d,
		rawaddr,
	}
}

func (i IPCAddress) String() string {
	// `-` is used as delimiter instead of `/`
	// `/` is harder to parse as `SubnetId` contains `/`, which makes it difficult to
	// determined which is the start of Address
	return strings.Join([]string{i.SubnetID.String(), i.RawAddress.String()}, IPCAddrSeparator)
}

func IPCAddressFromString(addr string) (IPCAddress, error) {
	r := strings.Split(addr, "-")
	if len(r) != 2 {
		return Undef, xerrors.New("invalid IPCAddress string type") // TODO Create new error, TODO Define Undef for IPCAddress
	}
	rawAddress, err := address.NewFromString(r[1])
	if err != nil {
		return Undef, err
	}
	subnetID, err := SubnetIDFromString(r[0])
	if err != nil {
		return Undef, err
	}
	return IPCAddress{subnetID, rawAddress}, nil
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
	IPCAddrSeparator = "-"
)

// RootSubnet is the ID of the root network
var RootSubnet = SubnetID{
	Parent: RootStr,
	Actor:  id0,
}

// UndefSubnetID is the undef ID
var UndefSubnetID = SubnetID{
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

func (id SubnetID) CommonParent(other SubnetID) (SubnetID, int) {
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
				return UndefSubnetID, 0
			}
			return sn, l
		}
	}
	sn, err := SubnetIDFromString(out)
	if err != nil {
		return UndefSubnetID, 0
	}
	return sn, l
}

func SubnetIDFromString(str string) (SubnetID, error) {
	switch str {
	case RootStr:
		return RootSubnet, nil
	case UndefStr:
		return UndefSubnetID, nil
	}

	s1 := strings.Split(str, SubnetSeparator)
	actor, err := address.NewFromString(s1[len(s1)-1]) // TODO Ask
	if err != nil {
		return UndefSubnetID, err
	}
	return SubnetID{
		Parent: strings.Join(s1[:len(s1)-1], SubnetSeparator),
		Actor:  actor,
	}, nil
}

func (id SubnetID) Down(curr SubnetID) SubnetID {
	s1 := strings.Split(id.String(), SubnetSeparator)
	s2 := strings.Split(curr.String(), SubnetSeparator)
	// curr needs to be contained in id
	if len(s2) >= len(s1) {
		return UndefSubnetID
	}
	_, l := id.CommonParent(curr)
	out := SubnetSeparator
	for i := 0; i <= l+1 && i < len(s1); i++ {
		if i < len(s2) && s1[i] != s2[i] {
			// they are not in a common path
			return UndefSubnetID
		}
		out = path.Join(out, s1[i])
	}
	sn, err := SubnetIDFromString(out)
	if err != nil {
		return UndefSubnetID
	}
	return sn
}

func (id SubnetID) Up(curr SubnetID) SubnetID {
	s1 := strings.Split(id.String(), SubnetSeparator)
	s2 := strings.Split(curr.String(), SubnetSeparator)
	// curr needs to be contained in id
	if len(s2) > len(s1) {
		return UndefSubnetID
	}

	_, l := id.CommonParent(curr)
	out := SubnetSeparator
	for i := 0; i < l; i++ {
		if i < len(s1) && s1[i] != s2[i] {
			// they are not in a common path
			return UndefSubnetID
		}
		out = path.Join(out, s1[i])
	}
	sn, err := SubnetIDFromString(out)
	if err != nil {
		return UndefSubnetID
	}
	return sn
}
