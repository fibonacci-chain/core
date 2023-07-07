package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fibonacci-chain/core/x/dex/exchange"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	"github.com/fibonacci-chain/core/x/dex/types"
)

func SetPriceStateFromExecutionOutcome(
	ctx sdk.Context,
	keeper *keeper.Keeper,
	contractAddr types.ContractAddress,
	pair types.Pair,
	outcome exchange.ExecutionOutcome,
) {
	if outcome.TotalQuantity.IsZero() {
		return
	}

	avgPrice := outcome.TotalNotional.Quo(outcome.TotalQuantity)
	priceState := types.Price{
		Pair:                       &pair,
		Price:                      avgPrice,
		SnapshotTimestampInSeconds: uint64(ctx.BlockTime().Unix()),
	}
	keeper.SetPriceState(ctx, priceState, string(contractAddr))
}
