package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/InnerPeace080/cosmos-test/testutil/keeper"
	"github.com/InnerPeace080/cosmos-test/x/cosmostest/keeper"
	"github.com/InnerPeace080/cosmos-test/x/cosmostest/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmostestKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
