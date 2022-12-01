package keeper_test

import (
	"context"
	"testing"

	keepertest "cist/testutil/keeper"
	"cist/x/cist/keeper"
	"cist/x/cist/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CistKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
