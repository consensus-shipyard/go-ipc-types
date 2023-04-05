package utils_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	"github.com/test-go/testify/require"

	"github.com/consensus-shipyard/go-ipc-types/voting"
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
	c, _ := cid.Parse("bafy2bzacecnamqgqmifpluoeldx7zzglxcljo6oja4vrmtj7432rphldpdmm2")
	params := voting.Voting{
		ExecutableEpochQueue: []abi.ChainEpoch{1},
		EpochVoteSubmission:  c,
		Ratio:                voting.Ratio{Num: 2, Denom: 3},
	}
	buf := new(bytes.Buffer)
	err := params.MarshalCBOR(buf)
	require.NoError(t, err)
	t.Log("===== Hex serialization =====")
	t.Log(hex.EncodeToString(buf.Bytes()))
}
