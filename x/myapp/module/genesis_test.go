package myapp_test

import (
	"testing"

	keepertest "github.com/kiarash-naderi/myapp/testutil/keeper"
	"github.com/kiarash-naderi/myapp/testutil/nullify"
	myapp "github.com/kiarash-naderi/myapp/x/myapp/module"
	"github.com/kiarash-naderi/myapp/x/myapp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyappKeeper(t)
	myapp.InitGenesis(ctx, k, genesisState)
	got := myapp.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
