package query_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/fibonacci-chain/core/testutil/keeper"
	"github.com/fibonacci-chain/core/x/dex/keeper/query"
	"github.com/fibonacci-chain/core/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestRegisteredContractQuery(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wrapper := query.KeeperWrapper{Keeper: keeper}
	wctx := sdk.WrapSDKContext(ctx)
	expectedContractInfo := types.ContractInfoV2{
		Creator:      keepertest.TestAccount,
		ContractAddr: keepertest.TestContract,
		CodeId:       1,
		RentBalance:  1000000,
	}
	err := keeper.SetContract(ctx, &types.ContractInfoV2{
		Creator:      keepertest.TestAccount,
		ContractAddr: keepertest.TestContract,
		CodeId:       1,
		RentBalance:  1000000,
	})
	require.NoError(t, err)

	request := types.QueryRegisteredContractRequest{
		ContractAddr: keepertest.TestContract,
	}
	expectedResponse := types.QueryRegisteredContractResponse{
		ContractInfo: &expectedContractInfo,
	}
	t.Run("Registered Contract query", func(t *testing.T) {
		response, err := wrapper.GetRegisteredContract(wctx, &request)
		require.NoError(t, err)
		require.Equal(t, expectedResponse, *response)
	})
}
