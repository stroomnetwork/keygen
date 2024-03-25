package keygen

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

type KeysOutput struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	Network    string `json:"network"`
}

func GenerateRandomKeys(params *chaincfg.Params) (*KeysOutput, error) {
	// Generate a random seed at the recommended length.
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		return nil, err
	}

	// Generate a new master node using the seed.
	key, err := hdkeychain.NewMaster(seed, params)
	if err != nil {
		return nil, err
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}
	if HasOddY(privateKey.PubKey()) {
		privateKey = NegatePrivateKey(privateKey)
	}
	output := &KeysOutput{
		PublicKey:  hex.EncodeToString(privateKey.PubKey().SerializeCompressed()),
		PrivateKey: hex.EncodeToString(privateKey.Serialize()),
		Network:    params.Name,
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
