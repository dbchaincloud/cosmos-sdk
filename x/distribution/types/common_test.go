package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dbchaincloud/tendermint/crypto"
	"github.com/dbchaincloud/tendermint/crypto/algo"
)

// nolint:deadcode,unused,varcheck
var (
	delPk1       = algo.GenPrivKey().PubKey()
	delPk2       = algo.GenPrivKey().PubKey()
	delPk3       = algo.GenPrivKey().PubKey()
	delAddr1     = sdk.AccAddress(delPk1.Address())
	delAddr2     = sdk.AccAddress(delPk2.Address())
	delAddr3     = sdk.AccAddress(delPk3.Address())
	emptyDelAddr sdk.AccAddress

	valPk1       = algo.GenPrivKey().PubKey()
	valPk2       = algo.GenPrivKey().PubKey()
	valPk3       = algo.GenPrivKey().PubKey()
	valAddr1     = sdk.ValAddress(valPk1.Address())
	valAddr2     = sdk.ValAddress(valPk2.Address())
	valAddr3     = sdk.ValAddress(valPk3.Address())
	emptyValAddr sdk.ValAddress

	emptyPubkey crypto.PubKey
)
