package query

import (
	"github.com/fibonacci-chain/core/x/dex/types"
)

var _ types.QueryServer = KeeperWrapper{}
