package verify

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/fibonacci-chain/core/testutil/processblock"
	"github.com/fibonacci-chain/core/utils"
	"github.com/stretchr/testify/require"
)

// assuming only `ufibo` will get distributed
func Allocation(t *testing.T, app *processblock.App, f BlockRunnable, _ []signing.Tx) BlockRunnable {
	return func() []uint32 {
		// fees collected in T-1 are allocated in T's BeginBlock, so we can simply
		// query fee collector's balance since this function is called between T-1
		// and T.
		feeCollector := app.AccountKeeper.GetModuleAccount(app.Ctx(), authtypes.FeeCollectorName)
		feesCollectedInt := app.BankKeeper.GetBalance(app.Ctx(), feeCollector.GetAddress(), "ufibo")
		feesCollected := sdk.NewDecCoinFromCoin(feesCollectedInt)

		prevProposer := sdk.ValAddress(app.DistrKeeper.GetPreviousProposerConsAddr(app.Ctx())).String()
		votedValidators := utils.Map(app.GetAllValidators(), func(v stakingtypes.Validator) string {
			return v.GetOperator().String()
		})
		expectedOutstandingRewards := getOutstandingRewards(app)

		baseProposerReward := app.DistrKeeper.GetBaseProposerReward(app.Ctx())
		bonusProposerReward := app.DistrKeeper.GetBonusProposerReward(app.Ctx())
		proposerMultiplier := baseProposerReward.Add(bonusProposerReward.MulTruncate(sdk.OneDec())) // in test, every val always signs
		proposerReward := sdk.DecCoin{
			Denom:  "ufibo",
			Amount: feesCollected.Amount.MulTruncate(proposerMultiplier),
		}
		expectedOutstandingRewards[prevProposer] = expectedOutstandingRewards[prevProposer].Add(proposerReward)

		communityTax := app.DistrKeeper.GetCommunityTax(app.Ctx())
		voteMultiplier := sdk.OneDec().Sub(proposerMultiplier).Sub(communityTax).QuoInt(sdk.NewInt(int64(len(votedValidators))))
		voterReward := sdk.DecCoin{
			Denom:  "ufibo",
			Amount: feesCollected.Amount.MulTruncate(voteMultiplier),
		}

		for _, validator := range votedValidators {
			expectedOutstandingRewards[validator] = expectedOutstandingRewards[validator].Add(voterReward)
		}

		res := f()

		actualOutstandingRewards := getOutstandingRewards(app)

		require.Equal(t, len(expectedOutstandingRewards), len(actualOutstandingRewards))

		for val, reward := range expectedOutstandingRewards {
			require.True(t, reward.Equal(actualOutstandingRewards[val]))
		}

		return res
	}
}

func getOutstandingRewards(app *processblock.App) map[string]sdk.DecCoin {
	outstandingRewards := map[string]sdk.DecCoin{}
	for _, val := range app.GetAllValidators() {
		outstandingRewards[val.GetOperator().String()] = sdk.NewDecCoin("ufibo", sdk.NewInt(0))
	}
	app.DistrKeeper.IterateValidatorOutstandingRewards(
		app.Ctx(),
		func(val sdk.ValAddress, rewards types.ValidatorOutstandingRewards) (stop bool) {
			if len(rewards.Rewards) == 0 {
				return false
			}
			if len(rewards.Rewards) > 1 {
				panic("expecting only ufibo as validator reward denom but found multiple")
			}
			if rewards.Rewards[0].Denom != "ufibo" {
				panic(fmt.Sprintf("expecting only ufibo as validator reward denom but found %s", rewards.Rewards[0].Denom))
			}
			outstandingRewards[val.String()] = rewards.Rewards[0]
			return false
		},
	)
	return outstandingRewards
}
