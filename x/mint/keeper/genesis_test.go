package keeper_test

import (
	"testing"
	"time"

	"github.com/fibonacci-chain/core/app"
	"github.com/fibonacci-chain/core/testutil/nullify"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/fibonacci-chain/core/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	now := time.Now()

	params := types.DefaultParams()
	params.TokenReleaseSchedule = []types.ScheduledTokenRelease{
		{
			StartDate:          now.Format(types.TokenReleaseDateFormat),
			EndDate:            now.Format(types.TokenReleaseDateFormat),
			TokenReleaseAmount: 100,
		},
	}
	genesisState := types.GenesisState{
		Params: params,
		Minter: types.Minter{
			StartDate:           now.Format(types.TokenReleaseDateFormat),
			EndDate:             now.Format(types.TokenReleaseDateFormat),
			Denom:               "ufibo",
			TotalMintAmount:     100,
			RemainingMintAmount: 0,
			LastMintAmount:      100,
			LastMintDate:        "2023-04-01",
			LastMintHeight:      0,
		},
	}

	app.MintKeeper.InitGenesis(ctx, &genesisState)
	got := app.MintKeeper.ExportGenesis(ctx)
	require.NotNil(t, got)
	require.Equal(t, got, &genesisState)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
