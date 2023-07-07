package dex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fibonacci-chain/core/x/dex/keeper"
	"github.com/fibonacci-chain/core/x/dex/types"
)

func HandleAddAssetMetadataProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddAssetMetadataProposal) error {
	for _, asset := range p.AssetList {
		k.SetAssetMetadata(ctx, asset)
	}
	return nil
}
