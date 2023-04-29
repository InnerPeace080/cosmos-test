package cosmostest_test

import (
	"testing"

	keepertest "github.com/InnerPeace080/cosmos-test/testutil/keeper"
	"github.com/InnerPeace080/cosmos-test/testutil/nullify"
	"github.com/InnerPeace080/cosmos-test/x/cosmostest"
	"github.com/InnerPeace080/cosmos-test/x/cosmostest/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 76,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmostestKeeper(t)
	cosmostest.InitGenesis(ctx, *k, genesisState)
	got := cosmostest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	// this line is used by starport scaffolding # genesis/test/assert
}
