package epoch

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fibonacci-chain/core/x/epoch/keeper"
	"github.com/fibonacci-chain/core/x/epoch/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetEpoch(
		ctx,
		*genState.Epoch,
	)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	epoch := k.GetEpoch(ctx)
	genesis.Epoch = &epoch

	return genesis
}
