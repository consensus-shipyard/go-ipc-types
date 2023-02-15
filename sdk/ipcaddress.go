package sdk

import (
	"fmt"
	"strings"

	"github.com/filecoin-project/go-address"
)

// UndefAddress creates a convenient type to define
// undefined IPC address
var UndefAddress = Address{}

// Address adds subnet information to raw Filecoin addresses
type Address struct {
	SubnetID   SubnetID
	RawAddress address.Address
}

func AddressFromString(addr string) (Address, error) {
	r := strings.Split(addr, IPCAddrSeparator)
	if len(r) != 2 {
		return UndefAddress, fmt.Errorf("invalid Address string type") // TODO Create new error, TODO Define Undef for Address
	}
	rawAddress, err := address.NewFromString(r[1])
	if err != nil {
		return UndefAddress, err
	}
	subnetID, err := NewSubnetIDFromString(r[0])
	if err != nil {
		return UndefAddress, err
	}
	return Address{subnetID, rawAddress}, nil
}
