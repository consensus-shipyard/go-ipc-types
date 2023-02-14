package subnetactor

import (
	"github.com/consensus-shipyard/go-ipc-types/ipcsdk"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type Validator struct {
	Addr    address.Address
	NetAddr string
}

type ConstructParams struct {
	Parent            ipcsdk.SubnetID
	Name              string
	IpcGatewayAddr    uint64
	Consensus         ConsensusType
	MinValidatorStake abi.TokenAmount
	MinValidators     uint64
	FinalityThreshold abi.ChainEpoch
	CheckPeriod       abi.ChainEpoch
	Genesis           []byte
}

type JoinParams struct {
	ValidatorNetAddr string
}

type Votes struct {
	Validators []address.Address
}

// ConsensusType defines the types of consensus supported
// by subnets.
type ConsensusType int64

const (
	Delegated ConsensusType = iota
	PoW
	Tendermint
	Mir
	FilecoinEC
	Dummy
)