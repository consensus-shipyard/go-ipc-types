package subnetactor

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/consensus-shipyard/go-ipc-types/sdk"
)

// ManifestID is the id used to index the gateway actor
// in the builtin-actors bundle.
const ManifestID = "ipc_subnet_actor"

type ConstructParams struct {
	Parent            sdk.SubnetID
	Name              string
	IPCGatewayAddr    uint64
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
type ConsensusType uint64

const (
	Delegated ConsensusType = iota
	PoW
	Tendermint
	Mir
	FilecoinEC
	Dummy
)
