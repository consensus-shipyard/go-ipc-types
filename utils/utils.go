package utils

import (
	"fmt"
	mbig "math/big"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/filecoin-project/specs-actors/v7/actors/builtin"
	"github.com/filecoin-project/specs-actors/v7/actors/util/adt"
)

func GetOutOfAdtArray[T any](adtArray *adt.Array, nonce uint64) (*T, error) {
	var (
		out   T
		found bool
		err   error
	)

	if i, ok := (any(&out)).(cbor.Unmarshaler); ok {
		found, err = adtArray.Get(nonce, i)
	} else {
		return nil, fmt.Errorf("the type *%T does not implement the cbor.Unmarshaler interface", out)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get %T with nonce %v: %w", out, nonce, err)
	}
	if !found {
		return nil, nil
	}
	return &out, nil
}

// GetOutOfArray takes a generic type that must implement cbor.Unmarshaler
// and returns a particular vale of a TAmt type passed as cid.Cid given the nonce
// If the type does not implement cbor.Unmarshaler then this returns an error at runtime
func GetOutOfArray[T any](cID cid.Cid, s adt.Store, nonce uint64, bitwidth int) (*T, error) {
	adtArray, err := adt.AsArray(s, cID, bitwidth)
	if err != nil {
		return nil, err
	}
	out, err := GetOutOfAdtArray[T](adtArray, nonce)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetOutOfHamt takes a generic type that must implement cbor.Unmarshaler
// and returns a particular value of a THamt type passed as cid.Cid given the key
// If the type does not implement cbor.Unmarshaler then this returns an error at runtime
func GetOutOfHamt[T any](cID cid.Cid, s adt.Store, k abi.Keyer) (*T, error) {
	var out T
	adtMap, err := adt.AsMap(s, cID, builtin.DefaultHamtBitwidth)
	if err != nil {
		return nil, fmt.Errorf("failed to get stake: %w", err)
	}

	if i, ok := (any(&out)).(cbor.Unmarshaler); ok {
		_, err = adtMap.Get(k, i)
	} else {
		return nil, fmt.Errorf("the type *%T does not implement the cbor.Unmarshaler interface", out)
	}
	return &out, err
}

func MajorityThreshold() mbig.Rat {
	x := mbig.NewInt(2)
	y := mbig.NewInt(3)
	return *new(mbig.Rat).SetFrac(x, y)
}
