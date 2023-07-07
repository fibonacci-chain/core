package utils

import (
	"context"

	dexcache "github.com/fibonacci-chain/core/x/dex/cache"
)

type MemStateKeyType string

const DexMemStateContextKey MemStateKeyType = MemStateKeyType("dex-memstate")

func GetMemState(ctx context.Context) *dexcache.MemState {
	if val := ctx.Value(DexMemStateContextKey); val != nil {
		return val.(*dexcache.MemState)
	}
	panic("cannot find mem state in context")
}
