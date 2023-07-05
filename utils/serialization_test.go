package utils_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/consensus-shipyard/go-ipc-types/sdk"
	"github.com/consensus-shipyard/go-ipc-types/subnetactor"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/test-go/testify/require"
)

// TestCborSerialization cbor-serializes a specific type and prints in a log
// the hex representation of the serialization.
//
// This is really helpful when debugging the serialization interop between Go
// and Rust, as the serialization can be input in tools like https://cbor.me/
// to inspect the result that was actually serialized.
// This is an example with subnetactor.ConstructParams, but it can be done
// with any type implementing the CborMarshaler interface.
func TestCborSerialization(t *testing.T) {
	// c, _ := cid.Parse("bafy2bzacecnamqgqmifpluoeldx7zzglxcljo6oja4vrmtj7432rphldpdmm2")
	id, _ := address.NewIDAddress(64)
	params := subnetactor.ConstructParams{
		Parent:              sdk.NewRootID(123),
		Name:                "test",
		IPCGatewayAddr:      id,
		BottomUpCheckPeriod: 30,
		TopDownCheckPeriod:  30,
		MinValidators:       1,
		MinValidatorStake:   abi.NewTokenAmount(100),
		Consensus:           subnetactor.Mir,
	}
	buf := new(bytes.Buffer)
	err := params.MarshalCBOR(buf)
	require.NoError(t, err)
	t.Log("===== Hex serialization =====")
	t.Log(hex.EncodeToString(buf.Bytes()))
}
