package contract

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	dexkeeperutils "github.com/fibonacci-chain/core/x/dex/keeper/utils"
	"github.com/fibonacci-chain/core/x/dex/types"
)

func HandleSettlements(
	ctx sdk.Context,
	contractAddr string,
	dexkeeper *keeper.Keeper,
	settlements []*types.SettlementEntry,
) error {
	return callSettlementHook(ctx, contractAddr, dexkeeper, settlements)
}

func callSettlementHook(
	ctx sdk.Context,
	contractAddr string,
	dexkeeper *keeper.Keeper,
	settlementEntries []*types.SettlementEntry,
) error {
	if len(settlementEntries) == 0 {
		return nil
	}
	_, currentEpoch := dexkeeper.IsNewEpoch(ctx)
	nativeSettlementMsg := types.SudoSettlementMsg{
		Settlement: types.Settlements{
			Epoch:   int64(currentEpoch),
			Entries: settlementEntries,
		},
	}
	if _, err := dexkeeperutils.CallContractSudo(ctx, dexkeeper, contractAddr, nativeSettlementMsg, dexkeeper.GetSettlementGasAllowance(ctx, len(settlementEntries))); err != nil {
		return err
	}
	return nil
}
