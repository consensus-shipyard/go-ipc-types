package types_test

import (
	"bytes"
	"fmt"
	"testing"

	ipct "github.com/consensus-shipyard/go-ipc-types/types"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-address"
)

func TestNaming(t *testing.T) {
	address.CurrentNetwork = address.Mainnet
	addr1, err := address.NewIDAddress(101)
	require.NoError(t, err)
	addr2, err := address.NewIDAddress(102)
	require.NoError(t, err)
	root := ipct.RootSubnet
	ptrNet1 := ipct.NewSubnetID(root, addr1)
	ptrNet2 := ipct.NewSubnetID(*ptrNet1, addr2)

	t.Log("Test actors")
	actor1 := ptrNet1.Actor
	require.Equal(t, actor1, addr1)
	actor2 := ptrNet2.Actor
	require.NoError(t, err)
	require.Equal(t, actor2, addr2)
	actorRoot := root.Actor
	require.NoError(t, err)
	id0, _ := address.NewIDAddress(0)
	require.Equal(t, id0, actorRoot)

	t.Log("Test parents")
	parent1 := ptrNet1.Parent
	require.NoError(t, err)
	require.Equal(t, root.String(), parent1)
	parent2 := ptrNet2.Parent
	require.NoError(t, err)
	require.Equal(t, parent2, ptrNet1.String())
	parentRoot := root.Parent
	require.NoError(t, err)
	require.Equal(t, parentRoot, ipct.RootStr)
}

func TestCborMarshal(t *testing.T) {
	addr1, err := address.NewIDAddress(101)
	require.NoError(t, err)
	root := ipct.RootSubnet
	net1 := ipct.NewSubnetID(root, addr1)

	var buf bytes.Buffer
	err = net1.MarshalCBOR(&buf)
	require.NoError(t, err)
	net2 := ipct.SubnetID{}
	err = net2.UnmarshalCBOR(&buf)
	require.NoError(t, err)
	require.Equal(t, net1, &net2)
}

func TestHAddress(t *testing.T) {
	address.CurrentNetwork = address.Mainnet
	id, _ := address.NewIDAddress(1000)
	a := ipct.IPCAddress{ipct.RootSubnet, id}

	sn := a.SubnetID
	require.Equal(t, ipct.RootSubnet, sn)

	raw := a.RawAddress
	require.Equal(t, id, raw)
}

func TestSubnetID(t *testing.T) {
	id, err := ipct.NewSubnetIDFromString("/root/f01")
	require.NoError(t, err)
	require.Equal(t, "/root/f01", id.String())
}

func TestSubnetOps(t *testing.T) {
	address.CurrentNetwork = address.Mainnet
	testParentAndBottomUp(t, "/root/f01", "/root/f01/f02", "/root/f01", 2, false)
	testParentAndBottomUp(t, "/root/f03/f01", "/root/f01/f02", "/root", 1, true)
	testParentAndBottomUp(t, "/root/f03/f01/f04", "/root/f03/f01/f05", "/root/f03/f01", 3, true)
	testParentAndBottomUp(t, "/root/f03/f01", "/root/f03/f02", "/root/f03", 2, true)

	testDownOrUp(t, "/root/f01/f02/f03", "/root/f01", "/root/f01/f02", true)
	testDownOrUp(t, "/root/f01/f02/f03", "/root/f01/f02", "/root/f01/f02/f03", true)
	testDownOrUp(t, "/root/f02", "/root/f01/f02/f03", ipct.UndefSubnetID.String(), true)
	testDownOrUp(t, "/root/f02", "/root/f02", ipct.UndefSubnetID.String(), true)

	testDownOrUp(t, "/root/f01/f02/f03", "/root/f01", "/root", false)
	testDownOrUp(t, "/root", "/root/f01", ipct.UndefSubnetID.String(), false)
	testDownOrUp(t, "/root/f01/f02/f03", "/root/f01/f02/f03/f05", ipct.UndefSubnetID.String(), false)
	testDownOrUp(t, "/root/f01/f02/f03", "/root/f01/f02", "/root/f01", false)
}

func testDownOrUp(t *testing.T, from, to, expected string, down bool) {
	sn, _ := ipct.SubnetIDFromString(from)
	arg, err := ipct.SubnetIDFromString(to)
	if err != nil {
		fmt.Println(err)
	}
	ex, _ := ipct.SubnetIDFromString(expected)
	if down {
		if expected != ipct.UndefSubnetID.String() {
			require.Equal(t, sn.Down(*arg), ex)
		} else {
			require.Equal(t, sn.Down(*arg) == nil, true)
		}
	} else {
		if expected != ipct.UndefSubnetID.String() {
			require.Equal(t, sn.Up(*arg), ex)
		} else {
			require.Equal(t, sn.Up(*arg) == nil, true)
		}
	}
}

func testParentAndBottomUp(t *testing.T, from, to, parent string, exl int, bottomup bool) {
	sFrom, err := ipct.SubnetIDFromString(from)
	require.NoError(t, err)
	sTo, err := ipct.SubnetIDFromString(to)
	require.NoError(t, err)
	p, l := sFrom.CommonParent(*sTo)
	sparent, err := ipct.SubnetIDFromString(parent)
	require.NoError(t, err)
	require.Equal(t, p, sparent)
	require.Equal(t, exl, l)
	require.Equal(t, ipct.IsBottomup(*sFrom, *sTo), bottomup)

}
