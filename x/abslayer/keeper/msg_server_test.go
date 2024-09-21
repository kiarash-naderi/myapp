
package keeper_test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/require"

    keepertest "github.com/kiarash-naderi/myapp/testutil/keeper"
    "github.com/kiarash-naderi/myapp/x/abslayer/keeper"
    "github.com/kiarash-naderi/myapp/x/abslayer/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
    k, ctx := keepertest.AbslayerKeeper(t)
    return keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
    ms, ctx := setupMsgServer(t)
    require.NotNil(t, ms)
    require.NotNil(t, ctx)
}