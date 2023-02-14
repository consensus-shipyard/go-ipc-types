package ipcsdk

import (
	"fmt"
	"strings"

	"github.com/filecoin-project/go-address"
)

// UndefIPCAddress creates a convenient type to define
// undefined IPC address
var UndefIPCAddress = IPCAddress{}

// IPCAddress adds subnet information to raw Filecoin addresses
type IPCAddress struct {
	SubnetID   SubnetID
	RawAddress address.Address
}

func IPCAddressFromString(addr string) (IPCAddress, error) {
	r := strings.Split(addr, IPCAddrSeparator)
	if len(r) != 2 {
		return UndefIPCAddress, fmt.Errorf("invalid IPCAddress string type") // TODO Create new error, TODO Define Undef for IPCAddress
	}
	rawAddress, err := address.NewFromString(r[1])
	if err != nil {
		return UndefIPCAddress, err
	}
	subnetID, err := NewSubnetIDFromString(r[0])
	if err != nil {
		return UndefIPCAddress, err
	}
	return IPCAddress{subnetID, rawAddress}, nil
}
