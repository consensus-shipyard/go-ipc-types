// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package validator

import (
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufValidator = []byte{130}

func (t *Validator) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufValidator); err != nil {
		return err
	}

	// t.Addr (address.Address) (struct)
	if err := t.Addr.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.NetAddr (string) (string)
	if len(t.NetAddr) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.NetAddr was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.NetAddr))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.NetAddr)); err != nil {
		return err
	}
	return nil
}

func (t *Validator) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Validator{}

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

	// t.Addr (address.Address) (struct)

	{

		if err := t.Addr.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Addr: %w", err)
		}

	}
	// t.NetAddr (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.NetAddr = string(sval)
	}
	return nil
}

var lengthBufSet = []byte{130}

func (t *Set) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufSet); err != nil {
		return err
	}

	// t.ConfigurationNumber (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.ConfigurationNumber)); err != nil {
		return err
	}

	// t.Validators ([]validator.Validator) (slice)
	if len(t.Validators) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Validators was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Validators))); err != nil {
		return err
	}
	for _, v := range t.Validators {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *Set) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Set{}

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

	// t.ConfigurationNumber (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ConfigurationNumber = uint64(extra)

	}
	// t.Validators ([]validator.Validator) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Validators: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Validators = make([]Validator, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v Validator
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Validators[i] = v
	}

	return nil
}
