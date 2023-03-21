package gateway

import (
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"

	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/utils"
)

type State struct {
	NetworkName          sdk.SubnetID
	TotalSubnets         uint64
	MinStake             abi.TokenAmount
	Subnets              cid.Cid // TCid<THamt<Cid, Subnet>>
	CheckPeriod          abi.ChainEpoch
	Checkpoints          cid.Cid // TCid<THamt<ChainEpoch, Checkpoint>>
	CheckMsgRegistry     cid.Cid // TCid<THamt<TCid<TLink<CrossMsgs>>, CrossMsgs>>
	Postbox              cid.Cid // TCid<THamt<Cid, Vec<u8>>>
	Nonce                uint64
	BottomupNonce        uint64
	BottomupMsgMeta      cid.Cid // TCid<TAmt<CrossMsgMeta, CROSSMSG_AMT_BITWIDTH>>
	AppliedBottomupNonce uint64
	AppliedTopdownNonce  uint64
}

func (st *State) GetSubnet(s adt.Store, id sdk.SubnetID) (*Subnet, bool, error) {
	return utils.GetOutOfHamt[Subnet](st.Subnets, s, id)
}

func (st *State) GetCheckpoints(s adt.Store, c abi.ChainEpoch) (*Checkpoint, bool, error) {
	return utils.GetOutOfHamt[Checkpoint](st.Checkpoints, s, abi.UIntKey(uint64(c)))
}

func (st *State) GetCrossMsgs(s adt.Store, cID cid.Cid) (*CrossMsgs, bool, error) {
	return utils.GetOutOfHamt[CrossMsgs](st.Checkpoints, s, abi.CidKey(cID))
}

func (st *State) GetBottomUpMsgMeta(s adt.Store, cID cid.Cid, nonce uint64) (*CrossMsgMeta, bool, error) {
	return utils.GetOutOfArray[CrossMsgMeta](cID, s, nonce, CrossMsgsAMTBitwidth)
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

// GetBottomUpMsgsFromRegistry returns the crossmsgs from a CID in the registry.
func (st *State) GetBottomUpMsgsFromRegistry(store adt.Store, c cid.Cid) (*CrossMsgs, bool, error) {
	msgMetas, err := adt.AsMap(store, st.CheckMsgRegistry, builtin.DefaultHamtBitwidth)
	if err != nil {
		return nil, false, err
	}
	var out CrossMsgs
	found, err := msgMetas.Get(abi.CidKey(c), &out)
	if err != nil {
		return nil, false, xerrors.Errorf("failed to get crossMsgMeta from registry with cid %v: %w", c, err)
	}
	if !found {
		return nil, false, nil
	}
	return &out, true, nil
}

// BottomUpMsgFromNonce gets the latest bottomUpMetas from a specific nonce
// (including the one specified, i.e. [nonce, latest], both limits
// included).
func (st *State) BottomUpMsgsFromNonce(s adt.Store, nonce uint64) ([]*CrossMsgMeta, error) {
	crossMsgs, err := adt.AsArray(s, st.BottomupMsgMeta, CrossMsgsAMTBitwidth)
	if err != nil {
		return nil, err
	}
	out := make([]*CrossMsgMeta, 0)
	for i := nonce; i < st.BottomupNonce; i++ {
		msg, found, err := getBottomUpMsg(crossMsgs, i)
		if err != nil {
			return nil, err
		}
		if found {
			out = append(out, msg)
		}
	}
	return out, nil
}

func getBottomUpMsg(crossMsgs *adt.Array, nonce uint64) (*CrossMsgMeta, bool, error) {
	var out CrossMsgMeta
	found, err := crossMsgs.Get(nonce, &out)
	if err != nil {
		return nil, false, xerrors.Errorf("failed to get cross-msg with nonce %v: %w", nonce, err)
	}
	if !found {
		return nil, found, nil
	}
	return &out, found, nil
}

// GetWindowCheckpoint gets the template for a specific epoch. If no template is persisted
// yet, an empty template is provided.
//
// NOTE: This function doesn't check if a template from the future is being requested.
func (st *State) GetWindowCheckpoint(s adt.Store, epoch abi.ChainEpoch) (*Checkpoint, error) {
	ch, found, err := utils.GetOutOfHamt[Checkpoint](st.Checkpoints, s,
		abi.UIntKey(uint64(CheckpointEpoch(epoch, st.CheckPeriod))))
	if err != nil {
		return nil, err
	}
	if !found {
		return NewCheckpoint(st.NetworkName, epoch), nil
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
