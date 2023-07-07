package keeper

import (
	"github.com/fibonacci-chain/core/x/epoch/types"
)

var _ types.QueryServer = Keeper{}
