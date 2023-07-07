package msgserver_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/fibonacci-chain/core/testutil/keeper"
	"github.com/fibonacci-chain/core/x/dex/keeper/msgserver"
	"github.com/fibonacci-chain/core/x/dex/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DexKeeper(t)
	return msgserver.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
