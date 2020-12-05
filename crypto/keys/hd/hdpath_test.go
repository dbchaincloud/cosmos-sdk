package hd

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/dbchaincloud/cosmos-sdk/types"

	bip39 "github.com/cosmos/go-bip39"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var defaultBIP39Passphrase = ""

// return bip39 seed with empty passphrase
func mnemonicToSeed(mnemonic string) []byte {
	return bip39.NewSeed(mnemonic, defaultBIP39Passphrase)
}

// nolint:govet
func ExampleStringifyPathParams() {
	path := NewParams(44, 0, 0, false, 0)
	fmt.Println(path.String())
	path = NewParams(44, 33, 7, true, 9)
	fmt.Println(path.String())
	// Output:
	// 44'/0'/0'/0/0
	// 44'/33'/7'/1/9
}

func TestStringifyFundraiserPathParams(t *testing.T) {
	path := NewFundraiserParams(4, types.CoinType, 22)
	require.Equal(t, "44'/118'/4'/0/22", path.String())

	path = NewFundraiserParams(4, types.CoinType, 57)
	require.Equal(t, "44'/118'/4'/0/57", path.String())

	path = NewFundraiserParams(4, 12345, 57)
	require.Equal(t, "44'/12345'/4'/0/57", path.String())
}

func TestPathToArray(t *testing.T) {
	path := NewParams(44, 118, 1, false, 4)
	require.Equal(t, "[44 118 1 0 4]", fmt.Sprintf("%v", path.DerivationPath()))

	path = NewParams(44, 118, 2, true, 15)
	require.Equal(t, "[44 118 2 1 15]", fmt.Sprintf("%v", path.DerivationPath()))
}

func TestParamsFromPath(t *testing.T) {
	goodCases := []struct {
		params *BIP44Params
		path   string
	}{
		{&BIP44Params{44, 0, 0, false, 0}, "44'/0'/0'/0/0"},
		{&BIP44Params{44, 1, 0, false, 0}, "44'/1'/0'/0/0"},
		{&BIP44Params{44, 0, 1, false, 0}, "44'/0'/1'/0/0"},
		{&BIP44Params{44, 0, 0, true, 0}, "44'/0'/0'/1/0"},
		{&BIP44Params{44, 0, 0, false, 1}, "44'/0'/0'/0/1"},
		{&BIP44Params{44, 1, 1, true, 1}, "44'/1'/1'/1/1"},
		{&BIP44Params{44, 118, 52, true, 41}, "44'/118'/52'/1/41"},
	}

	for i, c := range goodCases {
		params, err := NewParamsFromPath(c.path)
		errStr := fmt.Sprintf("%d %v", i, c)
		assert.NoError(t, err, errStr)
		assert.EqualValues(t, c.params, params, errStr)
		assert.Equal(t, c.path, c.params.String())
	}

	badCases := []struct {
		path string
	}{
		{"43'/0'/0'/0/0"},   // doesnt start with 44
		{"44'/1'/0'/0/0/5"}, // too many fields
		{"44'/0'/1'/0"},     // too few fields
		{"44'/0'/0'/2/0"},   // change field can only be 0/1
		{"44/0'/0'/0/0"},    // first field needs '
		{"44'/0/0'/0/0"},    // second field needs '
		{"44'/0'/0/0/0"},    // third field needs '
		{"44'/0'/0'/0'/0"},  // fourth field must not have '
		{"44'/0'/0'/0/0'"},  // fifth field must not have '
		{"44'/-1'/0'/0/0"},  // no negatives
		{"44'/0'/0'/-1/0"},  // no negatives
		{"a'/0'/0'/-1/0"},   // valid values
		{"0/X/0'/-1/0"},     // valid values
		{"44'/0'/X/-1/0"},   // valid values
		{"44'/0'/0'/%/0"},   // valid values
		{"44'/0'/0'/0/%"},   // valid values
	}

	for i, c := range badCases {
		params, err := NewParamsFromPath(c.path)
		errStr := fmt.Sprintf("%d %v", i, c)
		assert.Nil(t, params, errStr)
		assert.Error(t, err, errStr)
	}

}

// nolint:govet
func ExampleSomeBIP32TestVecs() {

	seed := mnemonicToSeed("barrel original fuel morning among eternal " +
		"filter ball stove pluck matrix mechanic")
	master, ch := ComputeMastersFromSeed(seed)
	fmt.Println("keys from fundraiser test-vector (cosmos, bitcoin, ether)")
	fmt.Println()
	// cosmos
	priv, err := DerivePrivateKeyForPath(master, ch, types.FullFundraiserPath)
	if err != nil {
		fmt.Println("INVALID")
	} else {
		fmt.Println(hex.EncodeToString(priv[:]))
	}
	// bitcoin
	priv, err = DerivePrivateKeyForPath(master, ch, "44'/0'/0'/0/0")
	if err != nil {
		fmt.Println("INVALID")
	} else {
		fmt.Println(hex.EncodeToString(priv[:]))
	}
	// ether
	priv, err = DerivePrivateKeyForPath(master, ch, "44'/60'/0'/0/0")
	if err != nil {
		fmt.Println("INVALID")
	} else {
		fmt.Println(hex.EncodeToString(priv[:]))
	}
	// INVALID
	priv, err = DerivePrivateKeyForPath(master, ch, "X/0'/0'/0/0")
	if err != nil {
		fmt.Println("INVALID")
	} else {
		fmt.Println(hex.EncodeToString(priv[:]))
	}
	priv, err = DerivePrivateKeyForPath(master, ch, "-44/0'/0'/0/0")
	if err != nil {
		fmt.Println("INVALID")
	} else {
		fmt.Println(hex.EncodeToString(priv[:]))
	}

	fmt.Println()
	fmt.Println("keys generated via https://coinomi.com/recovery-phrase-tool.html")
	fmt.Println()

	seed = mnemonicToSeed(
		"advice process birth april short trust crater change bacon monkey medal garment " +
			"gorilla ranch hour rival razor call lunar mention taste vacant woman sister")
	master, ch = ComputeMastersFromSeed(seed)
	priv, _ = DerivePrivateKeyForPath(master, ch, "44'/1'/1'/0/4")
	fmt.Println(hex.EncodeToString(priv[:]))

	seed = mnemonicToSeed("idea naive region square margin day captain habit " +
		"gun second farm pact pulse someone armed")
	master, ch = ComputeMastersFromSeed(seed)
	priv, _ = DerivePrivateKeyForPath(master, ch, "44'/0'/0'/0/420")
	fmt.Println(hex.EncodeToString(priv[:]))

	fmt.Println()
	fmt.Println("BIP 32 example")
	fmt.Println()

	// bip32 path: m/0/7
	seed = mnemonicToSeed("monitor flock loyal sick object grunt duty ride develop assault harsh history")
	master, ch = ComputeMastersFromSeed(seed)
	priv, _ = DerivePrivateKeyForPath(master, ch, "0/7")
	fmt.Println(hex.EncodeToString(priv[:]))

	// Output: keys from fundraiser test-vector (cosmos, bitcoin, ether)
	//
	// 4c2e449eb56b471cbd920b5c12b88367bded6660178082d91bafc5714e67da97
	// 6b267c9be47575ced3495f6a498c46a5e49b14a378cfd6e2a65dca7997566361
	// b2595f08ae281f076b32fa747b0b6bd99fe3f64e512f339e4742759a81df66ac
	// INVALID
	// INVALID
	//
	// keys generated via https://coinomi.com/recovery-phrase-tool.html
	//
	// 11b1e5c554f2a4595b2c7dc3b966c6b8ca450ef1dc13e039c8f6719046bffb35
	// b85e62de48af277a986a47c7d0310b13f5112d6bb320e2e243ddef428cd1e7b8
	//
	// BIP 32 example
	//
	// bfcf99c021e9db61b4d66ade727f23bf071f0ed082559336a956602f0cc08859
}
