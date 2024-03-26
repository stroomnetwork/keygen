package keygen

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomKeys(t *testing.T) {
	keys, err := GenerateRandomKeys()
	require.NoError(t, err)

	decoded, err := hex.DecodeString(keys.PrivateKey)
	require.NoError(t, err)
	privateKey, publicKey := btcec.PrivKeyFromBytes(decoded)
	require.False(t, HasOddY(publicKey))

	privateKeyBytes := privateKey.Serialize()
	publicKeyBytes := publicKey.SerializeCompressed()
	require.Equal(t, keys.PrivateKey, hex.EncodeToString(privateKeyBytes))
	require.Equal(t, keys.PublicKey, hex.EncodeToString(publicKeyBytes))
	require.Len(t, publicKeyBytes, 33)
	require.Len(t, privateKeyBytes, 32)
}
