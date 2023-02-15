package gateway

//go:generate go run ./gen/gen.go

import (
	"fmt"

	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/utils"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"
	"github.com/ipfs/go-cid"
)

type State struct {
	NetworkName          sdk.SubnetID
	TotalSubnets         uint64
	MinStake             abi.TokenAmount
	Subnets              cid.Cid // TCid<THamt<Cid, Subnet>>
	CheckPeriod          abi.ChainEpoch
	Checkpoints          cid.Cid //TCid<THamt<ChainEpoch, Checkpoint>>
	CheckMsgRegistry     cid.Cid //TCid<THamt<TCid<TLink<CrossMsgs>>, CrossMsgs>>
	Postbox              cid.Cid // TCid<THamt<Cid, Vec<u8>>>;
	Nonce                uint64
	BottomupNonce        uint64
	BottomupMsgMeta      cid.Cid // TCid<TAmt<CrossMsgMeta, CROSSMSG_AMT_BITWIDTH>>
	AppliedBottomupNonce uint64
	AppliedTopdownNonce  uint64
}

func GetTopDownMsg(crossMsgs *adt.Array, nonce uint64) (*CrossMsg, error) {
	var out CrossMsg
	found, err := crossMsgs.Get(nonce, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get cross-msg with nonce %v: %w", nonce, err)
	}
	if !found {
		return nil, nil
	}
	return &out, nil
}

func (st *State) GetSubnet(s adt.Store, id sdk.SubnetID) (*Subnet, error) {
	key, err := abi.ParseUIntKey(id.String())
	id.Bytes()
	if err != nil {
		return nil, err
	}
	subnet, err := utils.GetOutOfHamt[Subnet](st.Subnets, s, abi.UIntKey(key))
	return subnet, err
}

func (st *State) GetCheckpoints(s adt.Store, c abi.ChainEpoch) (*Checkpoint, error) {
	checkpoint, err := utils.GetOutOfHamt[Checkpoint](st.Checkpoints, s, abi.UIntKey(uint64(c)))
	return checkpoint, err
}

func (st *State) GetCrossMsgs(s adt.Store, cID cid.Cid) (*CrossMsgs, error) {
	crossMsgs, err := utils.GetOutOfHamt[CrossMsgs](st.Checkpoints, s, abi.CidKey(cID))
	return crossMsgs, err
}

func (st *State) GetBottomUpMsgMeta(s adt.Store, cID cid.Cid, nonce uint64) (*CrossMsgMeta, error) {
	return utils.GetOutOfArray[CrossMsgMeta](cID, s, nonce, CrossMsgsAMTBitwidth)
}

func (st *State) GetTopDownMsg(s adt.Store, id sdk.SubnetID, nonce uint64) (*CrossMsg, error) {
	sh, err := st.GetSubnet(s, id)
	if err != nil {
		return nil, err
	}
	CrossMsg, err := sh.GetTopDownMsg(s, nonce)
	return CrossMsg, err
}

// BottomUpMsgFromNonce gets the latest bottomUpMetas from a specific nonce
// (including the one specified, i.e. [nonce, latest], both limits
// included).
func (st *State) BottomUpMsgFromNonce(s adt.Store, nonce uint64) ([]*CrossMsgMeta, error) {
	out := make([]*CrossMsgMeta, 0)
	adtArray, err := adt.AsArray(s, st.BottomupMsgMeta, CrossMsgsAMTBitwidth)
	if err != nil {
		return nil, err
	}
	for i := nonce; i < st.BottomupNonce; i++ {
		meta, err := utils.GetOutOfAdtArray[CrossMsgMeta](adtArray, i)
		if err != nil {
			return nil, err
		}
		if meta != nil { // then found
			out = append(out, meta)
		}
	}
	return out, nil
}
