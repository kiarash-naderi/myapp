package keeper_test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/require"

    keepertest "github.com/kiarash-naderi/myapp/testutil/keeper"
    "github.com/kiarash-naderi/myapp/x/abslayer/keeper"
    "github.com/kiarash-naderi/myapp/x/abslayer/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServerWithKeeper(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
    k, ctx := keepertest.AbslayerKeeper(t)
    return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgUpdateParams(t *testing.T) {
    k, ms, ctx := setupMsgServerWithKeeper(t)
    params := types.DefaultParams()
    require.NoError(t, k.SetParams(ctx.(sdk.Context), params))
    wctx := sdk.UnwrapSDKContext(ctx)

    // default params
    testCases := []struct {
        name      string
        input     *types.MsgUpdateParams
        expErr    bool
        expErrMsg string
    }{
        {
            name: "invalid authority",
            input: &types.MsgUpdateParams{
                Authority: "invalid",
                Params:    params,
            },
            expErr:    true,
            expErrMsg: "invalid authority",
        },
        {
            name: "send enabled param",
            input: &types.MsgUpdateParams{
                Authority: k.GetAuthority(),
                Params:    types.Params{},
            },
            expErr: false,
        },
        {
            name: "all good",
            input: &types.MsgUpdateParams{
                Authority: k.GetAuthority(),
                Params:    params,
            },
            expErr: false,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            _, err := ms.UpdateParams(wctx, tc.input)

            if tc.expErr {
                require.Error(t, err)
                require.Contains(t, err.Error(), tc.expErrMsg)
            } else {
                require.NoError(t, err)
            }
        })
    }
}