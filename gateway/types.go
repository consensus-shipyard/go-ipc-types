package gateway

import (
	"github.com/consensus-shipyard/go-ipc-types/ipcsdk"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

type StorableMsg struct {
	From   ipcsdk.IPCAddress
	To     ipcsdk.IPCAddress
	Method abi.MethodNum
	Params RawBytes
	Value  abi.TokenAmount
	Nonce  uint64
}

func (sm *StorableMsg) IPCType() IPCMsgType {
	toSubnetID := sm.To.SubnetID
	fromSubnetID := sm.From.SubnetID
	if ipcsdk.IsBottomup(fromSubnetID, toSubnetID) {
		return IPCMsgTypeBottomUp
	}
	return IPCMsgTypeTopDown
}

const (
	IPCMsgTypeUnknown IPCMsgType = iota
	IPCMsgTypeBottomUp
	IPCMsgTypeTopDown
)

type IPCMsgType int

type ConstructParams struct {
	NetworkName      string
	CheckpointPeriod abi.ChainEpoch
}

type RawBytes struct {
	Bytes []byte
}

type CrossMsg struct {
	Msg     StorableMsg
	Wrapped bool
}

type FundParams struct {
	Value abi.TokenAmount
}

type CrossMsgParams struct {
	CrossMsg    CrossMsg
	Destination ipcsdk.SubnetID
}

type ApplyMsgParams struct {
	CrossMsg CrossMsg
}

type CrossMsgs struct {
	Msgs  []CrossMsg
	Metas []CrossMsgMeta
}

const CrossMsgsAMTBitwidth = 3

type Checkpoint struct {
	Data CheckData
	Sig  []byte
}

func (c *Checkpoint) CrossMsgMeta(from, to *ipcsdk.SubnetID) (*CrossMsgMeta, bool) {
	for i, m := range c.Data.CrossMsgs {
		if *from == m.From && *to == m.To {
			return &c.Data.CrossMsgs[i], true
		}
	}
	return nil, false
}

func (s *Checkpoint) CrossMsgMetaIndex(from, to *ipcsdk.SubnetID) (int, bool) {
	for i, m := range s.Data.CrossMsgs {
		if *from == m.From && *to == m.To {
			return i, true
		}
	}
	return 0, false
}

func CheckpointEpoch(epoch, period abi.ChainEpoch) abi.ChainEpoch {
	return (epoch / period) * period
}

func WindowEpoch(epoch, period abi.ChainEpoch) abi.ChainEpoch {
	ind := epoch / period
	return period * (ind + 1)
}

type CheckData struct {
	Source    ipcsdk.SubnetID
	TipSet    []byte
	Epoch     abi.ChainEpoch
	PrevCheck cid.Cid // TCid<TLink<Checkpoint>>
	Children  []ChildCheck
	CrossMsgs []CrossMsgMeta
}

type ChildCheck struct {
	Source ipcsdk.SubnetID
	Checks cid.Cid // Vec<TCid<TLink<Checkpoint>>>,
}

type CrossMsgMeta struct {
	From    ipcsdk.SubnetID
	To      ipcsdk.SubnetID
	MsgsCID cid.Cid // TCid<TLink<CrossMsgs>>,
	Nonce   uint64
	Value   abi.TokenAmount
}
