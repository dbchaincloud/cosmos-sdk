package types

import (
	"github.com/dbchaincloud/cosmos-sdk/codec"
)

// ModuleCdc is a generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}