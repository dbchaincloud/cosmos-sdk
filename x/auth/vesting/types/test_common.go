// nolint noalias
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dbchaincloud/tendermint/crypto"
	"github.com/dbchaincloud/tendermint/crypto/algo"
)

// NewTestMsg generates a test message
func NewTestMsg(addrs ...sdk.AccAddress) *sdk.TestMsg {
	return sdk.NewTestMsg(addrs...)
}

// NewTestCoins coins to more than cover the fee
func NewTestCoins() sdk.Coins {
	return sdk.Coins{
		sdk.NewInt64Coin("atom", 10000000),
	}
}

// KeyTestPubAddr generates a test key pair
func KeyTestPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	key := algo.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}
