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

// NewValidatorFromString creates a validator based on the string in `Addr:Weight@NetworkAddr` format.
//
// An examples of a validator string:
//   - t1wpixt5mihkj75lfhrnaa6v56n27epvlgwparujy:10@/ip4/127.0.0.1/tcp/10000/p2p/12D3KooWJhKBXvytYgPCAaiRtiNLJNSFG5jreKDu2jiVpJetzvVJ
//
// FIXME: Consider using json serde for this to support multiple multiaddr for validators.
func NewValidatorFromString(s string) (*Validator, error) {
	w := big.NewInt(0)

	parts := strings.Split(s, "@")
	if len(parts) != 2 {
		return nil, fmt.Errorf("failed to parse validators string")
	}
	idAndWeight := parts[0]
	netAddr := parts[1]
	parts = strings.Split(idAndWeight, ":")
	if len(parts) > 2 {
		return nil, fmt.Errorf("weight or ID are incorrect")
	}

	id := parts[0]
	if len(parts) == 2 {
		n, err := big.FromString(parts[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse weight: %w", err)
		}
		w = n
	}

	if id == "" {
		return nil, fmt.Errorf("empty address")
	}

	if netAddr == "" {
		return nil, fmt.Errorf("empty network address")
	}

	a, err := addr.NewFromString(id)
	if err != nil {
		return nil, err
	}
	ma, err := multiaddr.NewMultiaddr(netAddr)
	if err != nil {
		return nil, err
	}

	return NewValidatorWithWeight(a, ma.String(), &w), nil
}

type Validator struct {
	Addr addr.Address `json:"addr"`
	// FIXME: Consider using a multiaddr
	NetAddr string           `json:"net_addr,omitempty"`
	Weight  *abi.TokenAmount `json:"weight,omitempty"`
}

func NewValidator(a addr.Address, netAddr string) *Validator {
	w := abi.NewTokenAmount(0)
	return &Validator{Addr: a, NetAddr: netAddr, Weight: &w}
}

func NewValidatorWithWeight(a addr.Address, netAddr string, w *big.Int) *Validator {
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

// OnchainValidators information stored in the gateway actor
type OnChainValidators struct {
	Validators  Set
	TotalWeight abi.TokenAmount
}

func NewOnChainValidatorsFromSet(set *Set) OnChainValidators {
	weight := abi.NewTokenAmount(0)
	for _, v := range set.Validators {
		weight = big.Add(weight, *v.Weight)
	}
	return OnChainValidators{
		Validators:  *set,
		TotalWeight: weight,
	}
}
