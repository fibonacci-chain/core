package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	"github.com/fibonacci-chain/core/x/dex/types"
)

func V14ToV15(ctx sdk.Context, dexkeeper keeper.Keeper) error {
	// This isn't the cleanest migration since it could potentially revert any dex params we have changed
	// but we haven't, so we'll just do this.
	defaultParams := types.DefaultParams()
	dexkeeper.SetParams(ctx, defaultParams)
	return nil
}
