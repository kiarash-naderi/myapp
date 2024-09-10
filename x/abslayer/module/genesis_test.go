package abslayer_test

import (
	"testing"

	keepertest "github.com/kiarash-naderi/myapp/testutil/keeper"
	"github.com/kiarash-naderi/myapp/testutil/nullify"
	abslayer "github.com/kiarash-naderi/myapp/x/abslayer/module"
	"github.com/kiarash-naderi/myapp/x/abslayer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AbslayerKeeper(t)
	abslayer.InitGenesis(ctx, k, genesisState)
	got := abslayer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
