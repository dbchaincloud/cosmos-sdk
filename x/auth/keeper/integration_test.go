package keeper_test

import (
	abci "github.com/dbchaincloud/tendermint/abci/types"

	"github.com/dbchaincloud/cosmos-sdk/simapp"
	sdk "github.com/dbchaincloud/cosmos-sdk/types"
	authtypes "github.com/dbchaincloud/cosmos-sdk/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}
