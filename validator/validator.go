package validator

import (
	"fmt"
	"strings"

	"github.com/multiformats/go-multiaddr"
	"go.uber.org/zap/buffer"

	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)

// NewValidatorFromString parses a string containing of validator address and multiaddress separated by "@".
//
// Example of validator strings:
// - t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ
// FIXME: Consider using json serde for this to support multiple multiaddr for validators.
func NewValidatorFromString(s string) (*Validator, error) {
	parts := strings.Split(s, "@")
	if len(parts) != 2 {
		return nil, fmt.Errorf("incorrect format of the string")
	}

	a, err := addr.NewFromString(parts[0])
	if err != nil {
		return nil, err
	}
	m, err := multiaddr.NewMultiaddr(parts[1])
	if err != nil {
		return nil, err
	}

	return NewValidator(a, m.String()), nil
}

type Validator struct {
	Addr addr.Address `json:"addr"`
	// FIXME: Consider using a multiaddr
	NetAddr string          `json:"net_addr,omitempty"`
	Weight  abi.TokenAmount `json:"weight,omitempty"`
}

func NewValidator(a addr.Address, netAddr string) *Validator {
	return &Validator{Addr: a, NetAddr: netAddr, Weight: big.NewInt(0)}
}

func NewValidatorWithWeight(a addr.Address, netAddr string, w big.Int) *Validator {
	return &Validator{Addr: a, NetAddr: netAddr, Weight: w}
}

func (v *Validator) ID() string {
	return v.Addr.String()
}

func (v *Validator) Equal(o *Validator) bool {
	if v.Weight != o.Weight || v.Addr != o.Addr || v.NetAddr != o.NetAddr {
		return false
	}
	return true
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
