package gateway

import (
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"

	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/utils"
	"github.com/consensus-shipyard/go-ipc-types/validator"
	"github.com/consensus-shipyard/go-ipc-types/voting"
)

type State struct {
	NetworkName          sdk.SubnetID
	TotalSubnets         uint64
	MinStake             abi.TokenAmount
	Subnets              cid.Cid // TCid<THamt<SubnetID, Subnet>>
	BottomUpCheckPeriod  abi.ChainEpoch
	TopDownCheckPeriod   abi.ChainEpoch
	BottomUpCheckpoints  cid.Cid // TCid<THamt<ChainEpoch, Checkpoint>>
	Postbox              cid.Cid // TCid<THamt<Cid, Vec<u8>>>
	BottomupNonce        uint64
	AppliedBottomupNonce uint64
	AppliedTopdownNonce  uint64
	TopDownCheckVoting   voting.Voting
	Validators           validator.OnChainValidators
}

func (st *State) GetSubnet(s adt.Store, id sdk.SubnetID) (*Subnet, bool, error) {
	return utils.GetOutOfHamt[Subnet](st.Subnets, s, id)
}

func (st *State) GetCheckpoints(s adt.Store, e abi.ChainEpoch) (*BottomUpCheckpoint, bool, error) {
	return utils.GetOutOfHamt[BottomUpCheckpoint](st.BottomUpCheckpoints, s, sdk.EpochKey(e))
}

func (st *State) GetTopDownMsg(s adt.Store, id sdk.SubnetID, nonce uint64) (*CrossMsg, error) {
	sh, found, err := st.GetSubnet(s, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, xerrors.Errorf("subnet with id %s not found", id)
	}
	crossMsg, _, err := sh.GetTopDownMsg(s, nonce)
	return crossMsg, err
}

// GetWindowCheckpoint gets the template for a specific epoch. If no template is persisted
// yet, an empty template is provided.
//
// NOTE: This function doesn't check if a template from the future is being requested.
func (st *State) GetWindowCheckpoint(s adt.Store, epoch abi.ChainEpoch) (*BottomUpCheckpoint, error) {
	ch, found, err := utils.GetOutOfHamt[BottomUpCheckpoint](st.BottomUpCheckpoints, s,
		abi.UIntKey(uint64(CheckpointEpoch(epoch, st.BottomUpCheckPeriod))))
	if err != nil {
		return nil, err
	}
	if !found {
		return NewBottomUpCheckpoint(st.NetworkName, epoch), nil
	}
	return ch, nil
}

// ListSubnets lists subnets registered in the gateway
func (st *State) ListSubnets(s adt.Store) ([]Subnet, error) {
	subnetMap, err := adt.AsMap(s, st.Subnets, builtin.DefaultHamtBitwidth)
	if err != nil {
		return nil, err
	}

	var sh Subnet
	var out []Subnet
	err = subnetMap.ForEach(&sh, func(k string) error {
		out = append(out, sh)
		return nil
	})
	return out, err
}
