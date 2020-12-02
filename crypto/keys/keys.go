package keys

import "github.com/tendermint/tendermint/crypto/algo"

// SigningAlgo defines an algorithm to derive key-pairs which can be used for cryptographic signing.
type SigningAlgo string

const (
	// MultiAlgo implies that a pubkey is a multisignature
	MultiAlgo = SigningAlgo("multi")
	// Secp256k1 uses the Bitcoin secp256k1 ECDSA parameters.
	Secp256k1 = SigningAlgo("secp256k1")
	// Ed25519 represents the Ed25519 signature system.
	// It is currently not supported for end-user keys (wallets/ledgers).
	Ed25519 = SigningAlgo("ed25519")
	// Sr25519 represents the Sr25519 signature system.
	Sr25519 = SigningAlgo("sr25519")
	//sm2
	Sm2 = SigningAlgo("sm2")
)

//set global Algo type
var Algo = func()SigningAlgo{
	KeyType := algo.GetPubKeyType()
	switch KeyType {
	case algo.SM2:
		return Sm2
	case algo.ED25519:
		return Ed25519
	}
	return Secp256k1
}()
