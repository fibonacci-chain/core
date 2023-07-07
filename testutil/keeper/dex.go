package keeper

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/fibonacci-chain/core/app"
	dexcache "github.com/fibonacci-chain/core/x/dex/cache"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	"github.com/fibonacci-chain/core/x/dex/types"
	dexutils "github.com/fibonacci-chain/core/x/dex/utils"
	epochkeeper "github.com/fibonacci-chain/core/x/epoch/keeper"
	epochtypes "github.com/fibonacci-chain/core/x/epoch/types"
	minttypes "github.com/fibonacci-chain/core/x/mint/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

const (
	TestAccount    = "fb1yezq49upxhunjjhudql2fnj5dgvcwjj87pn2wx"
	TestContract   = "fb1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqladqwc"
	TestAccount2   = "fb1vk2f6aps83xahv2sql4equx8fa95jgcnsdxkvr"
	TestContract2  = "fb17p9rzwnnfxcjp32un9ug7yhhzgtkhvl9jfksztgw5uh69wac2pgsrtqewe"
	TestPriceDenom = "usdc"
	TestAssetDenom = "atom"
)

var (
	TestTicksize = sdk.OneDec()
	TestPair     = types.Pair{
		PriceDenom:       TestPriceDenom,
		AssetDenom:       TestAssetDenom,
		PriceTicksize:    &TestTicksize,
		QuantityTicksize: &TestTicksize,
	}
)

func TestApp() *app.App {
	return app.Setup(false)
}

func DexKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(typesparams.StoreKey)
	tKeyParams := sdk.NewTransientStoreKey(typesparams.TStoreKey)
	keyEpochs := sdk.NewKVStoreKey(epochtypes.StoreKey)
	dexMemStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(epochtypes.MemStoreKey)

	blackListAddrs := map[string]bool{}

	maccPerms := map[string][]string{
		types.ModuleName:     nil,
		minttypes.ModuleName: {authtypes.Minter},
	}

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyBank, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(dexMemStoreKey, sdk.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(keyEpochs, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tKeyParams, sdk.StoreTypeTransient, db)
	stateStore.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	cdc := codec.NewProtoCodec(app.MakeEncodingConfig().InterfaceRegistry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"DexParams",
	)
	paramsKeeper := paramskeeper.NewKeeper(cdc, codec.NewLegacyAmino(), keyParams, tKeyParams)
	accountKeeper := authkeeper.NewAccountKeeper(cdc, keyAcc, paramsKeeper.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms)
	bankKeeper := bankkeeper.NewBaseKeeper(cdc, keyBank, accountKeeper, paramsKeeper.Subspace(banktypes.ModuleName), blackListAddrs)
	epochKeeper := epochkeeper.NewKeeper(cdc, keyEpochs, memStoreKey, paramsKeeper.Subspace(epochtypes.ModuleName))
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		dexMemStoreKey,
		paramsSubspace,
		*epochKeeper,
		bankKeeper,
		accountKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	k.CreateModuleAccount(ctx)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())
	bankParams := banktypes.DefaultParams()
	bankParams.SendEnabled = []*banktypes.SendEnabled{
		{
			Denom:   TestPriceDenom,
			Enabled: true,
		},
	}
	bankKeeper.SetParams(ctx, bankParams)

	return k, ctx.WithContext(context.WithValue(ctx.Context(), dexutils.DexMemStateContextKey, dexcache.NewMemState(dexMemStoreKey)))
}

func CreateAssetMetadata(keeper *keeper.Keeper, ctx sdk.Context) types.AssetMetadata {
	ibcInfo := types.AssetIBCInfo{
		SourceChannel: "channel-1",
		DstChannel:    "channel-2",
		SourceDenom:   "uusdc",
		SourceChainID: "axelar",
	}

	denomUnit := banktypes.DenomUnit{
		Denom:    "ibc/D189335C6E4A68B513C10AB227BF1C1D38C746766278BA3EEB4FB14124F1D858",
		Exponent: 0,
		Aliases:  []string{"axlusdc", "usdc"},
	}

	var denomUnits []*banktypes.DenomUnit
	denomUnits = append(denomUnits, &denomUnit)

	metadata := banktypes.Metadata{
		Description: "Circle's stablecoin on Axelar",
		DenomUnits:  denomUnits,
		Base:        "ibc/D189335C6E4A68B513C10AB227BF1C1D38C746766278BA3EEB4FB14124F1D858",
		Name:        "USD Coin",
		Display:     "axlusdc",
		Symbol:      "USDC",
	}

	item := types.AssetMetadata{
		IbcInfo:   &ibcInfo,
		TypeAsset: "erc20",
		Metadata:  metadata,
	}

	keeper.SetAssetMetadata(ctx, item)

	return item
}

func SeedPriceSnapshot(ctx sdk.Context, k *keeper.Keeper, price string, timestamp uint64) {
	priceSnapshot := types.Price{
		SnapshotTimestampInSeconds: timestamp,
		Price:                      sdk.MustNewDecFromStr(price),
		Pair:                       &TestPair,
	}
	k.SetPriceState(ctx, priceSnapshot, TestContract)
}

func CreateNLongBook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LongBook {
	items := make([]types.LongBook, n)
	for i := range items {
		items[i].Entry = &types.OrderEntry{
			Price:      sdk.NewDec(int64(i)),
			Quantity:   sdk.NewDec(int64(i)),
			PriceDenom: TestPriceDenom,
			AssetDenom: TestAssetDenom,
		}
		items[i].Price = sdk.NewDec(int64(i))
		keeper.SetLongBook(ctx, TestContract, items[i])
	}
	return items
}

func CreateNShortBook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ShortBook {
	items := make([]types.ShortBook, n)
	for i := range items {
		items[i].Entry = &types.OrderEntry{
			Price:      sdk.NewDec(int64(i)),
			Quantity:   sdk.NewDec(int64(i)),
			PriceDenom: TestPriceDenom,
			AssetDenom: TestAssetDenom,
		}
		items[i].Price = sdk.NewDec(int64(i))
		keeper.SetShortBook(ctx, TestContract, items[i])
	}
	return items
}
