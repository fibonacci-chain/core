package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	"github.com/fibonacci-chain/core/x/dex/types"
)

func V11ToV12(ctx sdk.Context, dexkeeper keeper.Keeper) error {
	defaultParams := types.DefaultParams()
	dexkeeper.SetParams(ctx, defaultParams)
	return nil
}
