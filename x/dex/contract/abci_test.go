package contract_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	keepertest "github.com/fibonacci-chain/core/testutil/keeper"
	"github.com/fibonacci-chain/core/x/dex/contract"
	"github.com/fibonacci-chain/core/x/dex/types"
	minttypes "github.com/fibonacci-chain/core/x/mint/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestTransferRentFromDexToCollector(t *testing.T) {
	preRents := map[string]uint64{"abc": 100, "def": 50}
	postRents := map[string]uint64{"abc": 70}

	// expected total is (100 - 70) + 50 = 80
	testApp := keepertest.TestApp()
	ctx := testApp.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})
	bankkeeper := testApp.BankKeeper
	err := bankkeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin("ufibo", sdk.NewInt(100))))
	require.Nil(t, err)
	err = bankkeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ufibo", sdk.NewInt(100))))
	require.Nil(t, err)
	contract.TransferRentFromDexToCollector(ctx, bankkeeper, preRents, postRents)
	dexBalance := bankkeeper.GetBalance(ctx, testApp.AccountKeeper.GetModuleAddress(types.ModuleName), "ufibo")
	require.Equal(t, int64(20), dexBalance.Amount.Int64())
	collectorBalance := bankkeeper.GetBalance(ctx, testApp.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName), "ufibo")
	require.Equal(t, int64(80), collectorBalance.Amount.Int64())
}
