package migrations_test

import (
	"encoding/binary"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	"github.com/fibonacci-chain/core/x/dex/migrations"
	"github.com/fibonacci-chain/core/x/dex/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func TestMigrate4to5(t *testing.T) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"DexParams",
	)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	if !paramsSubspace.HasKeyTable() {
		paramsSubspace = paramsSubspace.WithKeyTable(types.ParamKeyTable())
	}
	store := ctx.KVStore(storeKey)
	store.Set([]byte("garbage key"), []byte("garbage value"))
	require.True(t, store.Has([]byte("garbage key")))
	err := migrations.V4ToV5(ctx, storeKey, paramsSubspace)
	require.Nil(t, err)
	require.False(t, store.Has([]byte("garbage key")))

	params := types.Params{}
	paramsSubspace.GetParamSet(ctx, &params)
	require.Equal(t, uint64(types.DefaultPriceSnapshotRetention), params.PriceSnapshotRetention)

	epochBytes := store.Get([]byte(keeper.EpochKey))
	epoch := binary.BigEndian.Uint64(epochBytes)
	require.Equal(t, uint64(0), epoch)
}
