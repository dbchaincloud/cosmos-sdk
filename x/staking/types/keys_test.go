package types

import (
	"encoding/hex"
	"github.com/tendermint/tendermint/crypto/algo"
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

var (
	keysPK1   = algo.GenPrivKeyFromSecret([]byte{1}).PubKey()
	keysPK2   = algo.GenPrivKeyFromSecret([]byte{2}).PubKey()
	keysPK3   = algo.GenPrivKeyFromSecret([]byte{3}).PubKey()
	keysAddr1 = keysPK1.Address()
	keysAddr2 = keysPK2.Address()
	keysAddr3 = keysPK3.Address()
)

func TestGetValidatorPowerRank(t *testing.T) {
	valAddr1 := sdk.ValAddress(keysAddr1)
	emptyDesc := Description{}
	val1 := NewValidator(valAddr1, keysPK1, emptyDesc)
	val1.Tokens = sdk.ZeroInt()
	val2, val3, val4 := val1, val1, val1
	val2.Tokens = sdk.TokensFromConsensusPower(1)
	val3.Tokens = sdk.TokensFromConsensusPower(10)
	x := new(big.Int).Exp(big.NewInt(2), big.NewInt(40), big.NewInt(0))
	val4.Tokens = sdk.TokensFromConsensusPower(x.Int64())

	tests := []struct {
		validator Validator
		wantHex   string
	}{
		{val1, "230000000000000000921b19ffe0d5cc11a7997e0fe49f9c00dbd709dd"},
		{val2, "230000000000000001921b19ffe0d5cc11a7997e0fe49f9c00dbd709dd"},
		{val3, "23000000000000000a921b19ffe0d5cc11a7997e0fe49f9c00dbd709dd"},
		{val4, "230000010000000000921b19ffe0d5cc11a7997e0fe49f9c00dbd709dd"},
	}
	for i, tt := range tests {
		got := hex.EncodeToString(getValidatorPowerRank(tt.validator))

		assert.Equal(t, tt.wantHex, got, "Keys did not match on test case %d", i)
	}
}

func TestGetREDByValDstIndexKey(t *testing.T) {
	tests := []struct {
		delAddr    sdk.AccAddress
		valSrcAddr sdk.ValAddress
		valDstAddr sdk.ValAddress
		wantHex    string
	}{
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr1),
			"366de4e6001f2a33ee586681f01b6063ff2428f6226de4e6001f2a33ee586681f01b6063ff2428f6226de4e6001f2a33ee586681f01b6063ff2428f622"},
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr2), sdk.ValAddress(keysAddr3),
			"367290ae649388e4836a47cb17f66842eb784e258d6de4e6001f2a33ee586681f01b6063ff2428f622ea6f3747a349f2ac37d026630f5f56d64010e979"},
		{sdk.AccAddress(keysAddr2), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr3),
			"367290ae649388e4836a47cb17f66842eb784e258dea6f3747a349f2ac37d026630f5f56d64010e9796de4e6001f2a33ee586681f01b6063ff2428f622"},
	}
	for i, tt := range tests {
		got := hex.EncodeToString(GetREDByValDstIndexKey(tt.delAddr, tt.valSrcAddr, tt.valDstAddr))

		assert.Equal(t, tt.wantHex, got, "Keys did not match on test case %d", i)
	}
}

func TestGetREDByValSrcIndexKey(t *testing.T) {
	tests := []struct {
		delAddr    sdk.AccAddress
		valSrcAddr sdk.ValAddress
		valDstAddr sdk.ValAddress
		wantHex    string
	}{
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr1),
			"356de4e6001f2a33ee586681f01b6063ff2428f6226de4e6001f2a33ee586681f01b6063ff2428f6226de4e6001f2a33ee586681f01b6063ff2428f622"},
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr2), sdk.ValAddress(keysAddr3),
			"35ea6f3747a349f2ac37d026630f5f56d64010e9796de4e6001f2a33ee586681f01b6063ff2428f6227290ae649388e4836a47cb17f66842eb784e258d"},
		{sdk.AccAddress(keysAddr2), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr3),
			"356de4e6001f2a33ee586681f01b6063ff2428f622ea6f3747a349f2ac37d026630f5f56d64010e9797290ae649388e4836a47cb17f66842eb784e258d"},
	}
	for i, tt := range tests {
		got := hex.EncodeToString(GetREDByValSrcIndexKey(tt.delAddr, tt.valSrcAddr, tt.valDstAddr))

		assert.Equal(t, tt.wantHex, got, "Keys did not match on test case %d", i)
	}
}
