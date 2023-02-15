package utils_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/consensus-shipyard/go-ipc-types/ipcsdk"
	"github.com/consensus-shipyard/go-ipc-types/subnetactor"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/test-go/testify/require"
)

// TestCborSerialization cbor-serializes a specific type and prints in a log
// the hex representation of the serialization.
//
// This is really helpful when debuggin the serialization interop between Go
// and Rust, as the serialization can be input in tools like https://cbor.me/
// to inspect the result that was actually serialized.
// This is an example with subnetactor.ConstructParams, but it can be done
// with any type implementing the CborMarshaler interface.
func TestCborSerialization(t *testing.T) {
	params := subnetactor.ConstructParams{
		Parent:            ipcsdk.RootSubnet,
		Name:              "test",
		IpcGatewayAddr:    64,
		CheckPeriod:       0,
		FinalityThreshold: 0,
		MinValidators:     0,
		MinValidatorStake: big.Zero(),
		Consensus:         subnetactor.Mir,
	}
	buf := new(bytes.Buffer)
	err := params.MarshalCBOR(buf)
	require.NoError(t, err)
	t.Log("===== Hex serialization =====")
	t.Log(hex.EncodeToString(buf.Bytes()))
}
