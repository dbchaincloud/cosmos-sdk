package exported_test

import (
	"github.com/dbchaincloud/tendermint/crypto/algo"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisAccountsContains(t *testing.T) {
	pubkey := algo.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	acc := authtypes.NewBaseAccount(addr, nil, algo.GenPrivKey().PubKey(), 0, 0)

	genAccounts := exported.GenesisAccounts{}
	require.False(t, genAccounts.Contains(acc.GetAddress()))

	genAccounts = append(genAccounts, acc)
	require.True(t, genAccounts.Contains(acc.GetAddress()))
}
