package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dbchaincloud/tendermint/crypto/algo"
	"testing"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/dbchaincloud/cosmos-sdk/types"
	authexported "github.com/dbchaincloud/cosmos-sdk/x/auth/exported"
	authtypes "github.com/dbchaincloud/cosmos-sdk/x/auth/types"

	"github.com/stretchr/testify/require"
)

func TestModuleAccountMarshalYAML(t *testing.T) {
	name := "test"
	moduleAcc := NewEmptyModuleAccount(name, Minter, Burner, Staking)
	bs, err := yaml.Marshal(moduleAcc)
	require.NoError(t, err)

	want := "|\n  address: cosmos12hsjayt9p5h7c4hvwnsa8exahl8zauaxg4mjk4\n  coins: []\n  public_key: \"\"\n  account_number: 0\n  sequence: 0\n  name: test\n  permissions:\n  - minter\n  - burner\n  - staking\n"
	require.Equal(t, want, string(bs))
}

func TestHasPermissions(t *testing.T) {
	name := "test"
	macc := NewEmptyModuleAccount(name, Staking, Minter, Burner)
	cases := []struct {
		permission string
		expectHas  bool
	}{
		{Staking, true},
		{Minter, true},
		{Burner, true},
		{"other", false},
	}

	for i, tc := range cases {
		hasPerm := macc.HasPermission(tc.permission)
		if tc.expectHas {
			require.True(t, hasPerm, "test case #%d", i)
		} else {
			require.False(t, hasPerm, "test case #%d", i)
		}
	}
}

func TestValidate(t *testing.T) {
	addr := sdk.AccAddress(algo.GenPrivKey().PubKey().Address())
	baseAcc := authtypes.NewBaseAccount(addr, sdk.Coins{}, nil, 0, 0)
	tests := []struct {
		name   string
		acc    authexported.GenesisAccount
		expErr error
	}{
		{
			"valid module account",
			NewEmptyModuleAccount("test"),
			nil,
		},
		{
			"invalid name and address pair",
			NewModuleAccount(baseAcc, "test"),
			fmt.Errorf("address %s cannot be derived from the module name 'test'", addr),
		},
		{
			"empty module account name",
			NewModuleAccount(baseAcc, "    "),
			errors.New("module account name cannot be blank"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.acc.Validate()
			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestModuleAccountJSON(t *testing.T) {
	pubkey := algo.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	coins := sdk.NewCoins(sdk.NewInt64Coin("test", 5))
	baseAcc := authtypes.NewBaseAccount(addr, coins, nil, 10, 50)
	acc := NewModuleAccount(baseAcc, "test", "burner")

	bz, err := json.Marshal(acc)
	require.NoError(t, err)

	bz1, err := acc.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, string(bz1), string(bz))

	var a ModuleAccount
	require.NoError(t, json.Unmarshal(bz, &a))
	require.Equal(t, acc.String(), a.String())
}
