package sdk

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/filecoin-project/go-address"
)

type SubnetID struct {
	Parent string
	Actor  address.Address
}

var id0, _ = address.NewIDAddress(0)

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

func (id SubnetID) Key() string {
	return id.String()
}

func NewSubnetIDFromString(addr string) (SubnetID, error) {
	var out SubnetID
	if addr == RootSubnet.String() {
		out = RootSubnet
		return out, nil
	}
	dir, file := filepath.Split(addr)
	act, err := address.NewFromString(file)
	if err != nil {
		return UndefSubnetID, err
	}
	return SubnetID{
		Parent: dir[:len(dir)-1], // move trailing `/`
		Actor:  act,
	}, nil
}

func NewSubnetID(parent SubnetID, subnetAct address.Address) SubnetID {
	return SubnetID{
		parent.String(),
		subnetAct,
	}
}

func (id SubnetID) Bytes() []byte {
	return []byte(id.String())
}

// String returns the id in string form.
func (id SubnetID) String() string {
	if id == RootSubnet {
		if id.Parent != UndefStr {
			return RootStr
		}
		return UndefStr
	}
	return filepath.Join(id.Parent, id.Actor.String())
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
			sn, err := NewSubnetIDFromString(out)
			if err != nil {
				return UndefSubnetID, 0
			}
			return sn, l
		}
	}
	sn, err := NewSubnetIDFromString(out)
	if err != nil {
		return UndefSubnetID, 0
	}
	return sn, l
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
	sn, err := NewSubnetIDFromString(out)
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
	sn, err := NewSubnetIDFromString(out)
	if err != nil {
		return UndefSubnetID
	}
	return sn
}

func IsBottomUp(from SubnetID, to SubnetID) bool {
	subnetID, index := from.CommonParent(to)
	if subnetID == UndefSubnetID {
		return false
	}

	a := from.String()
	components := strings.Split(a, SubnetSeparator)
	count := len(components) - 1
	return count > index

}
