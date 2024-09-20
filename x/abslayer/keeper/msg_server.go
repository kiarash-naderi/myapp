package keeper

import (
    "context"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/kiarash-naderi/myapp/x/abslayer/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

type msgServer struct {
    Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
    return &msgServer{Keeper: keeper}
}

func (m msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
    if msg.Authority != m.GetAuthority() {
        return nil, sdkerrors.ErrUnauthorized.Wrap("invalid authority")
    }

    m.Keeper.SetParams(sdk.UnwrapSDKContext(ctx), msg.Params)
    return &types.MsgUpdateParamsResponse{}, nil
}