package validator

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-address"
)

func TestHashesAreEqualForEqualMemberships(t *testing.T) {
	v1, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	v2, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10001/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)

	vs1 := NewValidatorSet(0, []*Validator{v1, v2})
	h1, err := vs1.Hash()
	require.NoError(t, err)
	vs2 := NewValidatorSet(0, []*Validator{v2, v1})
	h2, err := vs2.Hash()
	require.NoError(t, err)
	require.Equal(t, h1, h2)

	vs3 := NewValidatorSet(1, []*Validator{v1, v2})
	h3, err := vs3.Hash()
	require.NoError(t, err)
	require.NotEqual(t, h1, h3)

}

type validatorStrTest struct {
	addr    string
	netAddr string
	weight  string
	correct bool
}

var validatorStrTests = []validatorStrTest{
	{
		"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy",
		"/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
		"1",
		true,
	},
	{
		"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy",
		"/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
		"1844674407370955161518446744073709551615",
		true,
	},
	{
		"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy",
		"/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
		"0",
		true,
	},
	{"",
		"/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
		"0",
		false,
	},
	{"",
		"/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
		"10",
		false,
	},
	{"",
		"/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
		"",
		false,
	},
	{
		"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy", "",
		"0",
		false,
	},
	{
		"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy", "",
		"20",
		false,
	},
	{
		"", "",
		"",
		false,
	},
}

var incorrectValidatorStrings = []string{
	"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy:t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
	"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
	"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
	"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy::/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
	"t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy:8",
	"8:/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ",
	"8:@",
}

func TestValidatorFromString(t *testing.T) {
	for _, test := range validatorStrTests {
		v, err := NewValidatorFromString(fmt.Sprintf("%v:%v@%v", test.addr, test.weight, test.netAddr))
		if !test.correct {
			require.Error(t, err)
			continue
		}

		require.Equal(t, test.addr, v.Addr.String())
		require.Equal(t, test.netAddr, v.NetAddr)
		require.Equal(t, test.weight, v.Weight.String())
	}

	for _, test := range incorrectValidatorStrings {
		v, err := NewValidatorFromString(test)
		require.Error(t, err)
		if v != nil {
			t.Fatalf("not nil validator ")
		}
	}
}

func TestValidatorSetFromEnv(t *testing.T) {
	nonce := "1"
	addr := "t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy"
	netAddr := "/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ"

	err := os.Setenv("TEST_VALIDATORS_FROM_ENV", fmt.Sprintf("%v;%v@%v", nonce, addr, netAddr))
	require.NoError(t, err)
	defer func() {
		err := os.Unsetenv("TEST_VALIDATORS_FROM_ENV")
		require.NoError(t, err)
	}()

	vs, err := NewValidatorSetFromEnv("TEST_VALIDATORS_FROM_ENV")
	require.NoError(t, err)

	a, err := address.NewFromString(vs.Validators[0].ID())
	require.NoError(t, err)
	require.Equal(t, a, vs.Validators[0].Addr)

	m, err := multiaddr.NewMultiaddr(vs.Validators[0].NetAddr)
	require.NoError(t, err)
	require.Equal(t, m.String(), vs.Validators[0].NetAddr)

}

func TestMembership(t *testing.T) {
	v1, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	v2, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10001/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	v3, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10002/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)

	v := []*Validator{v1, v2, v3}
	vs := NewValidatorSet(2, v)
	require.NoError(t, err)
	require.Equal(t, len(v), len(vs.Validators))
}

func TestValidatorSetFromFile(t *testing.T) {
	fileName := "_vs_test_file.tmp"
	t.Cleanup(func() {
		err := os.Remove(fileName)
		require.NoError(t, err)
	})

	v1, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	v2, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10001/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	v3, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10002/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)

	vs1 := NewValidatorSet(0, []*Validator{v1, v2, v3})
	require.Equal(t, 3, vs1.Size())
	require.Equal(t, uint64(0), vs1.GetConfigurationNumber())

	err = vs1.Save(fileName)
	require.NoError(t, err)

	vs2, err := NewValidatorSetFromFile(fileName)
	require.NoError(t, err)

	require.Equal(t, true, vs1.Equal(vs2))

	vs3 := NewValidatorSetFromValidators(0, []*Validator{v1, v2, v3}...)
	require.Equal(t, true, vs1.Equal(vs3))
}

func TestValidatorSetFromString(t *testing.T) {
	s1 := "t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ"
	s2 := "t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10001/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ"
	s3 := "t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10002/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ"

	v1, err := NewValidatorFromString(s1)
	require.NoError(t, err)
	v2, err := NewValidatorFromString(s2)
	require.NoError(t, err)
	v3, err := NewValidatorFromString(s3)
	require.NoError(t, err)

	vs1 := NewValidatorSet(3, []*Validator{v1, v2, v3})
	require.Equal(t, 3, vs1.Size())

	vs2, err := NewValidatorSetFromString("3;" + s1 + "," + s2 + "," + s3)
	require.NoError(t, err)

	require.Equal(t, true, vs1.Equal(vs2))
}

func TestValidatorSetFromJson(t *testing.T) {
	s1 := "t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10001/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ"
	s2 := "t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10002/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ"

	v1, err := NewValidatorFromString(s1)
	require.NoError(t, err)
	v2, err := NewValidatorFromString(s2)
	require.NoError(t, err)

	vs2 := NewValidatorSet(2, []*Validator{v1, v2})
	js := vs2.JSONString()

	var vs1 Set
	err = json.Unmarshal([]byte(js), &vs1)
	require.NoError(t, err)

	require.Equal(t, true, vs1.Equal(vs2))
}

func TestAddValidatorToFile(t *testing.T) {
	fileName := "_v_test_file_.tmp"
	t.Cleanup(func() {
		err := os.Remove(fileName)
		require.NoError(t, err)
	})

	v1, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	err = AddValidatorToFile(fileName, v1)
	require.NoError(t, err)

	vs1 := NewValidatorSet(0, []*Validator{v1})
	vs, err := NewValidatorSetFromFile(fileName)
	require.NoError(t, err)
	require.Equal(t, true, vs1.Equal(vs))

	v2, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10001/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	err = AddValidatorToFile(fileName, v2)
	require.NoError(t, err)

	v3, err := NewValidatorFromString("t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10002/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ")
	require.NoError(t, err)
	err = AddValidatorToFile(fileName, v3)
	require.NoError(t, err)

	vs1 = NewValidatorSet(2, []*Validator{v1, v2, v3})
	require.Equal(t, 3, vs1.Size())

	vs2, err := NewValidatorSetFromFile(fileName)
	require.NoError(t, err)

	require.Equal(t, true, vs1.Equal(vs2))
}
