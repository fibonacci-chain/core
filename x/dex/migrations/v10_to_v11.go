package migrations

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	"github.com/fibonacci-chain/core/x/dex/types"
)

var DexPrefixes = []string{
	types.LongBookKey,
	types.ShortBookKey,
	types.TriggerBookKey,
	types.OrderKey,
	types.TwapKey,
	types.RegisteredPairKey,
	types.OrderKey,
	types.CancelKey,
	types.AccountActiveOrdersKey,
	types.NextOrderIDKey,
	types.MatchResultKey,
	types.MemOrderKey,
	types.MemCancelKey,
	types.MemDepositKey,
	types.PriceKey,
	types.SettlementEntryKey,
	types.NextSettlementIDKey,
}

func V10ToV11(ctx sdk.Context, dexkeeper keeper.Keeper) error {
	dexkeeper.CreateModuleAccount(ctx)

	// this will nuke all old prefixes data in the store
	for _, prefixKey := range DexPrefixes {
		store := prefix.NewStore(ctx.KVStore(dexkeeper.GetStoreKey()), []byte(prefixKey))
		iterator := sdk.KVStorePrefixIterator(store, []byte{})

		defer iterator.Close()
		for ; iterator.Valid(); iterator.Next() {
			store.Delete(iterator.Key())
		}
	}

	return nil
}
