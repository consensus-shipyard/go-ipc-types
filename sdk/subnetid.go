package sdk

import (
	"path"
	"strconv"
	"strings"

	"github.com/filecoin-project/go-address"
)

type SubnetID struct {
	// ChainID of the root network
	Root uint64
	// Children up to the current subnet.
	Children []address.Address
}

const (
	RootPrefix       = "r"
	SubnetSeparator  = "/"
	UndefStr         = ""
	IPCAddrSeparator = ":"
)

// UndefSubnetID is the undef ID
var UndefSubnetID = SubnetID{
	Root:     0,
	Children: []address.Address{},
}

func (id SubnetID) Key() string {
	return id.String()
}

// Equal checks if two subnet IDs are equal.
func (id SubnetID) Equal(other SubnetID) bool {
	if id.Root != other.Root {
		return false
	}
	if len(id.Children) != len(other.Children) {
		return false
	}
	for i, c := range id.Children {
		if c != other.Children[i] {
			return false
		}
	}
	return true
}

func (id SubnetID) Parent() SubnetID {
	if len(id.Children) == 0 {
		return UndefSubnetID
	}
	return SubnetID{
		Root:     id.Root,
		Children: id.Children[:len(id.Children)-1],
	}
}

// Actor returns the subnet actor ID that governs the current subnet
// or the ID=0 if the current subnet is the root.
func (id SubnetID) Actor() address.Address {
	if len(id.Children) == 0 {
		id, _ := address.NewIDAddress(0)
		return id
	}
	return id.Children[len(id.Children)-1]
}

// NewRootID creates a new root subnet ID from its chainID.
func NewRootID(chainID uint64) SubnetID {
	return SubnetID{
		Root:     chainID,
		Children: []address.Address{},
	}
}

func NewSubnetIDFromString(addr string) (SubnetID, error) {
	cs := strings.Split(addr, "/")[1:]
	id, err := strconv.ParseUint(cs[0][1:], 10, 64)
	if err != nil {
		return UndefSubnetID, err
	}

	children := make([]address.Address, len(cs)-1)
	if len(cs) > 1 {
		for i, c := range cs[1:] {
			children[i], err = address.NewFromString(c)
			if err != nil {
				return UndefSubnetID, err
			}
		}
	}

	return SubnetID{
		Root:     id,
		Children: children,
	}, nil

}

func NewSubnetIDFromRoute(rootID uint64, route []address.Address) SubnetID {
	return SubnetID{
		Root:     rootID,
		Children: route,
	}
}

func NewSubnetID(parent SubnetID, subnetAct address.Address) SubnetID {
	return SubnetID{
		parent.Root,
		append(parent.Children, subnetAct),
	}
}

func (id SubnetID) Bytes() []byte {
	return []byte(id.String())
}

// String returns the id in string form.
func (id SubnetID) String() string {
	out := SubnetSeparator + RootPrefix + strconv.FormatUint(id.Root, 10) + SubnetSeparator
	for _, c := range id.Children {
		out = path.Join(out, c.String())
	}
	return out
}

// CommonParent computes the common parent of two subnet IDs and
// returns the number of children in it.
func (id SubnetID) CommonParent(other SubnetID) (SubnetID, int) {
	if id.Root != other.Root {
		return UndefSubnetID, 0
	}
	var i int
	for i = 0; i < len(id.Children) && i < len(other.Children); i++ {
		if id.Children[i] != other.Children[i] {
			break
		}
	}
	return SubnetID{
		Root:     id.Root,
		Children: id.Children[:i],
	}, i
}

// Down Returns from the current subnet the next subnet down in the path
// defined by the current subnet and the destination subnet.
func (id SubnetID) Down(curr SubnetID) SubnetID {
	if len(id.Children) <= len(curr.Children) {
		return UndefSubnetID
	}

	if cp, i := id.CommonParent(curr); !cp.Equal(UndefSubnetID) {
		return SubnetID{
			Root:     id.Root,
			Children: id.Children[:i+1],
		}
	}

	return UndefSubnetID
}

// Up returns the SubnetID immediately up from the current subnet.
func (id SubnetID) Up(curr SubnetID) SubnetID {
	if len(id.Children) < len(curr.Children) {
		return UndefSubnetID
	}

	if cp, i := id.CommonParent(curr); !cp.Equal(UndefSubnetID) {
		return SubnetID{
			Root:     id.Root,
			Children: id.Children[:i-1],
		}
	}

	return UndefSubnetID
}

// IsBottomUp returns true if the from subnet is above in the hierarchy.
func IsBottomUp(from SubnetID, to SubnetID) bool {
	_, i := to.CommonParent(from)
	return len(from.Children) > i
}
