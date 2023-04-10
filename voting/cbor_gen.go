// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package voting

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufVoting = []byte{134}

func (t *Voting) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufVoting); err != nil {
		return err
	}

	// t.GenesisEpoch (abi.ChainEpoch) (int64)
	if t.GenesisEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.GenesisEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.GenesisEpoch-1)); err != nil {
			return err
		}
	}

	// t.SubmissionPeriod (abi.ChainEpoch) (int64)
	if t.SubmissionPeriod >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SubmissionPeriod)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.SubmissionPeriod-1)); err != nil {
			return err
		}
	}

	// t.LastVotingExecuted (abi.ChainEpoch) (int64)
	if t.LastVotingExecuted >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.LastVotingExecuted)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.LastVotingExecuted-1)); err != nil {
			return err
		}
	}

	// t.ExecutableEpochQueue ([]abi.ChainEpoch) (slice)
	if len(t.ExecutableEpochQueue) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.ExecutableEpochQueue was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.ExecutableEpochQueue))); err != nil {
		return err
	}
	for _, v := range t.ExecutableEpochQueue {
		if v >= 0 {
			if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(v)); err != nil {
				return err
			}
		} else {
			if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-v-1)); err != nil {
				return err
			}
		}
	}

	// t.EpochVoteSubmission (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.EpochVoteSubmission); err != nil {
		return xerrors.Errorf("failed to write cid field t.EpochVoteSubmission: %w", err)
	}

	// t.Ratio (voting.Ratio) (struct)
	if err := t.Ratio.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *Voting) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Voting{}

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

	// t.GenesisEpoch (abi.ChainEpoch) (int64)
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

		t.GenesisEpoch = abi.ChainEpoch(extraI)
	}
	// t.SubmissionPeriod (abi.ChainEpoch) (int64)
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

		t.SubmissionPeriod = abi.ChainEpoch(extraI)
	}
	// t.LastVotingExecuted (abi.ChainEpoch) (int64)
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

		t.LastVotingExecuted = abi.ChainEpoch(extraI)
	}
	// t.ExecutableEpochQueue ([]abi.ChainEpoch) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.ExecutableEpochQueue: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.ExecutableEpochQueue = make([]abi.ChainEpoch, extra)
	}

	for i := 0; i < int(extra); i++ {
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

			t.ExecutableEpochQueue[i] = abi.ChainEpoch(extraI)
		}
	}

	// t.EpochVoteSubmission (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.EpochVoteSubmission: %w", err)
		}

		t.EpochVoteSubmission = c

	}
	// t.Ratio (voting.Ratio) (struct)

	{

		if err := t.Ratio.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Ratio: %w", err)
		}

	}
	return nil
}

var lengthBufRatio = []byte{130}

func (t *Ratio) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufRatio); err != nil {
		return err
	}

	// t.Num (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Num)); err != nil {
		return err
	}

	// t.Denom (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Denom)); err != nil {
		return err
	}

	return nil
}

func (t *Ratio) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Ratio{}

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

	// t.Num (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Num = uint64(extra)

	}
	// t.Denom (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Denom = uint64(extra)

	}
	return nil
}
