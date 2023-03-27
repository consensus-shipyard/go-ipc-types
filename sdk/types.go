package sdk

import (
	"encoding/binary"

	"github.com/filecoin-project/go-state-types/abi"
)

// Status defines the different states in which a subnet can be.
type Status int64

const (
	Instantiated Status = iota
	Active
	Inactive
	Terminating
	Killed
)

type EpochKey abi.ChainEpoch

func (k EpochKey) Key() string {
	// varInt integer encoding
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(k))
	return string(buf[:n])
}
