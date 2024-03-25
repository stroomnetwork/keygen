package keygen

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
)

type KeysOutput struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// GenerateRandomKeys generates `bip-0340` compatible keys and encodes in hex.
func GenerateRandomKeys() (*KeysOutput, error) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}
	if HasOddY(privateKey.PubKey()) {
		privateKey = NegatePrivateKey(privateKey)
	}
	output := &KeysOutput{
		PublicKey:  hex.EncodeToString(privateKey.PubKey().SerializeCompressed()),
		PrivateKey: hex.EncodeToString(privateKey.Serialize()),
	}

	return output, nil
}

// HasOddY returns true if y coordinate of pubKey is odd.
// BIP-340 requires that y coordinate is even.
// Does not change its arguments.
func HasOddY(pubKey *btcec.PublicKey) bool {
	return pubKey.Y().Bit(0) == 1
}

// NegatePrivateKey returns negated private key.
// For private key `a`, it returns `-a`. Such as `a + (-a) = 0`.
// Does not change its arguments.
func NegatePrivateKey(privateKey *btcec.PrivateKey) *btcec.PrivateKey {
	var negated btcec.ModNScalar
	negated.Set(&privateKey.Key)
	negated.Negate()

	var negatedPrivateKey btcec.PrivateKey
	negatedPrivateKey.Key = negated

	return &negatedPrivateKey
}
