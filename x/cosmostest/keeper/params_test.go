package keeper_test

import (
	"testing"

	testkeeper "github.com/InnerPeace080/cosmos-test/testutil/keeper"
	"github.com/InnerPeace080/cosmos-test/x/cosmostest/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmostestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
