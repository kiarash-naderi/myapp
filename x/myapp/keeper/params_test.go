package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/kiarash-naderi/myapp/testutil/keeper"
	"github.com/kiarash-naderi/myapp/x/myapp/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MyappKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
