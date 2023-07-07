package query

import (
	"github.com/fibonacci-chain/core/x/dex/keeper"
)

type KeeperWrapper struct {
	*keeper.Keeper
}
