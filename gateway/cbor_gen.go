// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package gateway

import (
	"fmt"
	"io"
	"math"
	"sort"

	sdk "github.com/consensus-shipyard/go-ipc-types/sdk"
	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufState = []byte{141}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufState); err != nil {
		return err
	}

	// t.NetworkName (sdk.SubnetID) (struct)
	if err := t.NetworkName.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TotalSubnets (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.TotalSubnets)); err != nil {
		return err
	}

	// t.MinStake (big.Int) (struct)
	if err := t.MinStake.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Subnets (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.Subnets); err != nil {
		return xerrors.Errorf("failed to write cid field t.Subnets: %w", err)
	}

	// t.BottomUpCheckPeriod (abi.ChainEpoch) (int64)
	if t.BottomUpCheckPeriod >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.BottomUpCheckPeriod)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.BottomUpCheckPeriod-1)); err != nil {
			return err
		}
	}

	// t.TopDownCheckPeriod (abi.ChainEpoch) (int64)
	if t.TopDownCheckPeriod >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.TopDownCheckPeriod)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.TopDownCheckPeriod-1)); err != nil {
			return err
		}
	}

	// t.BottomUpCheckpoints (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.BottomUpCheckpoints); err != nil {
		return xerrors.Errorf("failed to write cid field t.BottomUpCheckpoints: %w", err)
	}

	// t.Postbox (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.Postbox); err != nil {
		return xerrors.Errorf("failed to write cid field t.Postbox: %w", err)
	}

	// t.BottomupNonce (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.BottomupNonce)); err != nil {
		return err
	}

	// t.AppliedBottomupNonce (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.AppliedBottomupNonce)); err != nil {
		return err
	}

	// t.AppliedTopdownNonce (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.AppliedTopdownNonce)); err != nil {
		return err
	}

	// t.TopDownCheckVoting (voting.Voting) (struct)
	if err := t.TopDownCheckVoting.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Validators (validator.OnChainValidators) (struct)
	if err := t.Validators.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) (err error) {
	*t = State{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 13 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.NetworkName (sdk.SubnetID) (struct)

	{

		if err := t.NetworkName.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.NetworkName: %w", err)
		}

	}
	// t.TotalSubnets (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.TotalSubnets = uint64(extra)

	}
	// t.MinStake (big.Int) (struct)

	{

		if err := t.MinStake.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.MinStake: %w", err)
		}

	}
	// t.Subnets (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Subnets: %w", err)
		}

		t.Subnets = c

	}
	// t.BottomUpCheckPeriod (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.BottomUpCheckPeriod = abi.ChainEpoch(extraI)
	}
	// t.TopDownCheckPeriod (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.TopDownCheckPeriod = abi.ChainEpoch(extraI)
	}
	// t.BottomUpCheckpoints (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.BottomUpCheckpoints: %w", err)
		}

		t.BottomUpCheckpoints = c

	}
	// t.Postbox (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Postbox: %w", err)
		}

		t.Postbox = c

	}
	// t.BottomupNonce (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.BottomupNonce = uint64(extra)

	}
	// t.AppliedBottomupNonce (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.AppliedBottomupNonce = uint64(extra)

	}
	// t.AppliedTopdownNonce (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.AppliedTopdownNonce = uint64(extra)

	}
	// t.TopDownCheckVoting (voting.Voting) (struct)

	{

		if err := t.TopDownCheckVoting.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.TopDownCheckVoting: %w", err)
		}

	}
	// t.Validators (validator.OnChainValidators) (struct)

	{

		if err := t.Validators.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Validators: %w", err)
		}

	}
	return nil
}

var lengthBufConstructParams = []byte{130}

func (t *ConstructParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufConstructParams); err != nil {
		return err
	}

	// t.NetworkName (string) (string)
	if len(t.NetworkName) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.NetworkName was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.NetworkName))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.NetworkName)); err != nil {
		return err
	}

	// t.CheckpointPeriod (abi.ChainEpoch) (int64)
	if t.CheckpointPeriod >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.CheckpointPeriod)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.CheckpointPeriod-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *ConstructParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ConstructParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.NetworkName (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.NetworkName = string(sval)
	}
	// t.CheckpointPeriod (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.CheckpointPeriod = abi.ChainEpoch(extraI)
	}
	return nil
}

var lengthBufBottomUpCheckpoint = []byte{130}

func (t *BottomUpCheckpoint) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufBottomUpCheckpoint); err != nil {
		return err
	}

	// t.Data (gateway.CheckData) (struct)
	if err := t.Data.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Sig ([]uint8) (slice)
	if len(t.Sig) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Sig was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Sig))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Sig[:]); err != nil {
		return err
	}
	return nil
}

func (t *BottomUpCheckpoint) UnmarshalCBOR(r io.Reader) (err error) {
	*t = BottomUpCheckpoint{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Data (gateway.CheckData) (struct)

	{

		if err := t.Data.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Data: %w", err)
		}

	}
	// t.Sig ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Sig: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Sig = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Sig[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufTopDownCheckpoint = []byte{130}

func (t *TopDownCheckpoint) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufTopDownCheckpoint); err != nil {
		return err
	}

	// t.Epoch (abi.ChainEpoch) (int64)
	if t.Epoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Epoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Epoch-1)); err != nil {
			return err
		}
	}

	// t.TopDownMsgs ([]gateway.CrossMsgs) (slice)
	if len(t.TopDownMsgs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.TopDownMsgs was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.TopDownMsgs))); err != nil {
		return err
	}
	for _, v := range t.TopDownMsgs {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *TopDownCheckpoint) UnmarshalCBOR(r io.Reader) (err error) {
	*t = TopDownCheckpoint{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Epoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Epoch = abi.ChainEpoch(extraI)
	}
	// t.TopDownMsgs ([]gateway.CrossMsgs) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.TopDownMsgs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.TopDownMsgs = make([]CrossMsgs, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v CrossMsgs
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.TopDownMsgs[i] = v
	}

	return nil
}

var lengthBufCheckData = []byte{134}

func (t *CheckData) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCheckData); err != nil {
		return err
	}

	// t.Source (sdk.SubnetID) (struct)
	if err := t.Source.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Proof ([]uint8) (slice)
	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Proof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Proof[:]); err != nil {
		return err
	}

	// t.Epoch (abi.ChainEpoch) (int64)
	if t.Epoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Epoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Epoch-1)); err != nil {
			return err
		}
	}

	// t.PrevCheck (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.PrevCheck); err != nil {
		return xerrors.Errorf("failed to write cid field t.PrevCheck: %w", err)
	}

	// t.Children ([]gateway.ChildCheck) (slice)
	if len(t.Children) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Children was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Children))); err != nil {
		return err
	}
	for _, v := range t.Children {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}

	// t.CrossMsgs (gateway.BatchCrossMsgs) (struct)
	if err := t.CrossMsgs.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *CheckData) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CheckData{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Source (sdk.SubnetID) (struct)

	{

		if err := t.Source.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Source: %w", err)
		}

	}
	// t.Proof ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Proof = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Proof[:]); err != nil {
		return err
	}
	// t.Epoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Epoch = abi.ChainEpoch(extraI)
	}
	// t.PrevCheck (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PrevCheck: %w", err)
		}

		t.PrevCheck = c

	}
	// t.Children ([]gateway.ChildCheck) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Children: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Children = make([]ChildCheck, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v ChildCheck
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Children[i] = v
	}

	// t.CrossMsgs (gateway.BatchCrossMsgs) (struct)

	{

		if err := t.CrossMsgs.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.CrossMsgs: %w", err)
		}

	}
	return nil
}

var lengthBufChildCheck = []byte{130}

func (t *ChildCheck) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufChildCheck); err != nil {
		return err
	}

	// t.Source (sdk.SubnetID) (struct)
	if err := t.Source.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Checks ([]cid.Cid) (slice)
	if len(t.Checks) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Checks was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Checks))); err != nil {
		return err
	}
	for _, v := range t.Checks {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Checks: %w", err)
		}
	}
	return nil
}

func (t *ChildCheck) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ChildCheck{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Source (sdk.SubnetID) (struct)

	{

		if err := t.Source.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Source: %w", err)
		}

	}
	// t.Checks ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Checks: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Checks = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Checks failed: %w", err)
		}
		t.Checks[i] = c
	}

	return nil
}

var lengthBufBatchCrossMsgs = []byte{130}

func (t *BatchCrossMsgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufBatchCrossMsgs); err != nil {
		return err
	}

	// t.CrossMsgs ([]gateway.CrossMsg) (slice)
	if len(t.CrossMsgs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.CrossMsgs was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.CrossMsgs))); err != nil {
		return err
	}
	for _, v := range t.CrossMsgs {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}

	// t.Fee (big.Int) (struct)
	if err := t.Fee.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *BatchCrossMsgs) UnmarshalCBOR(r io.Reader) (err error) {
	*t = BatchCrossMsgs{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.CrossMsgs ([]gateway.CrossMsg) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.CrossMsgs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.CrossMsgs = make([]CrossMsg, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v CrossMsg
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.CrossMsgs[i] = v
	}

	// t.Fee (big.Int) (struct)

	{

		if err := t.Fee.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Fee: %w", err)
		}

	}
	return nil
}

var lengthBufStorableMsg = []byte{134}

func (t *StorableMsg) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufStorableMsg); err != nil {
		return err
	}

	// t.From (sdk.IPCAddress) (struct)
	if err := t.From.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.To (sdk.IPCAddress) (struct)
	if err := t.To.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Method (abi.MethodNum) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Method)); err != nil {
		return err
	}

	// t.Params ([]uint8) (slice)
	if len(t.Params) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Params was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Params))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Params[:]); err != nil {
		return err
	}

	// t.Value (big.Int) (struct)
	if err := t.Value.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	return nil
}

func (t *StorableMsg) UnmarshalCBOR(r io.Reader) (err error) {
	*t = StorableMsg{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.From (sdk.IPCAddress) (struct)

	{

		if err := t.From.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.From: %w", err)
		}

	}
	// t.To (sdk.IPCAddress) (struct)

	{

		if err := t.To.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.To: %w", err)
		}

	}
	// t.Method (abi.MethodNum) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Method = abi.MethodNum(extra)

	}
	// t.Params ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Params: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Params = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Params[:]); err != nil {
		return err
	}
	// t.Value (big.Int) (struct)

	{

		if err := t.Value.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Value: %w", err)
		}

	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	return nil
}

var lengthBufSubnet = []byte{135}

func (t *Subnet) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufSubnet); err != nil {
		return err
	}

	// t.ID (sdk.SubnetID) (struct)
	if err := t.ID.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Stake (big.Int) (struct)
	if err := t.Stake.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TopDownMsgs (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.TopDownMsgs); err != nil {
		return xerrors.Errorf("failed to write cid field t.TopDownMsgs: %w", err)
	}

	// t.Nonce (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	// t.CircSupply (big.Int) (struct)
	if err := t.CircSupply.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Status (sdk.Status) (int64)
	if t.Status >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Status-1)); err != nil {
			return err
		}
	}

	// t.PrevCheckpoint (gateway.BottomUpCheckpoint) (struct)
	if err := t.PrevCheckpoint.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *Subnet) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Subnet{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 7 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ID (sdk.SubnetID) (struct)

	{

		if err := t.ID.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ID: %w", err)
		}

	}
	// t.Stake (big.Int) (struct)

	{

		if err := t.Stake.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Stake: %w", err)
		}

	}
	// t.TopDownMsgs (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.TopDownMsgs: %w", err)
		}

		t.TopDownMsgs = c

	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	// t.CircSupply (big.Int) (struct)

	{

		if err := t.CircSupply.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.CircSupply: %w", err)
		}

	}
	// t.Status (sdk.Status) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Status = sdk.Status(extraI)
	}
	// t.PrevCheckpoint (gateway.BottomUpCheckpoint) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}
			t.PrevCheckpoint = new(BottomUpCheckpoint)
			if err := t.PrevCheckpoint.UnmarshalCBOR(cr); err != nil {
				return xerrors.Errorf("unmarshaling t.PrevCheckpoint pointer: %w", err)
			}
		}

	}
	return nil
}

var lengthBufCrossMsg = []byte{130}

func (t *CrossMsg) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCrossMsg); err != nil {
		return err
	}

	// t.Msg (gateway.StorableMsg) (struct)
	if err := t.Msg.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Wrapped (bool) (bool)
	if err := cbg.WriteBool(w, t.Wrapped); err != nil {
		return err
	}
	return nil
}

func (t *CrossMsg) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CrossMsg{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Msg (gateway.StorableMsg) (struct)

	{

		if err := t.Msg.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Msg: %w", err)
		}

	}
	// t.Wrapped (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.Wrapped = false
	case 21:
		t.Wrapped = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

var lengthBufFundParams = []byte{129}

func (t *FundParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufFundParams); err != nil {
		return err
	}

	// t.Value (big.Int) (struct)
	if err := t.Value.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *FundParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = FundParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Value (big.Int) (struct)

	{

		if err := t.Value.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Value: %w", err)
		}

	}
	return nil
}

var lengthBufCrossMsgParams = []byte{130}

func (t *CrossMsgParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCrossMsgParams); err != nil {
		return err
	}

	// t.CrossMsg (gateway.CrossMsg) (struct)
	if err := t.CrossMsg.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Destination (sdk.SubnetID) (struct)
	if err := t.Destination.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *CrossMsgParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CrossMsgParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.CrossMsg (gateway.CrossMsg) (struct)

	{

		if err := t.CrossMsg.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.CrossMsg: %w", err)
		}

	}
	// t.Destination (sdk.SubnetID) (struct)

	{

		if err := t.Destination.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Destination: %w", err)
		}

	}
	return nil
}

var lengthBufCrossMsgs = []byte{129}

func (t *CrossMsgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCrossMsgs); err != nil {
		return err
	}

	// t.Msgs ([]gateway.CrossMsg) (slice)
	if len(t.Msgs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Msgs was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Msgs))); err != nil {
		return err
	}
	for _, v := range t.Msgs {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *CrossMsgs) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CrossMsgs{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Msgs ([]gateway.CrossMsg) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Msgs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Msgs = make([]CrossMsg, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v CrossMsg
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Msgs[i] = v
	}

	return nil
}
