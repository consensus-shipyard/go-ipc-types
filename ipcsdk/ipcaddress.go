package ipcsdk

import (
	"fmt"
	"strings"

	"github.com/filecoin-project/go-address"
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
