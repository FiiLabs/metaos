package wservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/FiiLabs/metaos/modules/wservice/keeper"
)

func NewHandler(keeper keeper.IKeeper) sdk.Handler {
	return nil
}
