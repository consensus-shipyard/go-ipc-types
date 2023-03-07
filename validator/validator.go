package validator

import (
	"fmt"
	"strings"

	"github.com/multiformats/go-multiaddr"
	"go.uber.org/zap/buffer"

	addr "github.com/filecoin-project/go-address"
)

// NewValidatorFromString parses a string containing of validator address and multiaddress separated by "@".
//
// Example of validator strings:
// - t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ
// FIXME: Consider using json serde for this to support multiple multiaddr for validators.
func NewValidatorFromString(s string) (Validator, error) {
	parts := strings.Split(s, "@")
	if len(parts) != 2 {
		return Validator{}, fmt.Errorf("incorrect format of the string")
	}

	a, err := addr.NewFromString(parts[0])
	if err != nil {
		return Validator{}, err
	}
	m, err := multiaddr.NewMultiaddr(parts[1])
	if err != nil {
		return Validator{}, err
	}

	return Validator{
		Addr:    a,
		NetAddr: m.String(),
	}, nil
}

type Validator struct {
	Addr addr.Address `json:"addr"`
	// FIXME: Consider using a multiaddr
	NetAddr string `json:"net_addr,omitempty"`
	Weight  uint64 `json:"weight"`
}

func (v *Validator) ID() string {
	return v.Addr.String()
}

func (v *Validator) Bytes() ([]byte, error) {
	var b buffer.Buffer
	if err := v.MarshalCBOR(&b); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (v *Validator) String() string {
	return fmt.Sprintf("%s@%s", v.Addr.String(), v.NetAddr)
}

func SplitAndTrimEmpty(s, sep, cutset string) []string {
	if s == "" {
		return []string{}
	}

	spl := strings.Split(s, sep)
	nonEmptyStrings := make([]string, 0, len(spl))

	for i := 0; i < len(spl); i++ {
		element := strings.Trim(spl[i], cutset)
		if element != "" {
			nonEmptyStrings = append(nonEmptyStrings, element)
		}
	}

	return nonEmptyStrings
}
