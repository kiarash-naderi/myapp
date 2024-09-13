package keeper

import (
    "context"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/kiarash-naderi/myapp/x/abslayer/types"
)

type msgServer struct {
    Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
    return &msgServer{Keeper: keeper}
}

// SendToken handles MsgSendToken
func (m msgServer) SendToken(goCtx context.Context, msg *types.MsgSendToken) (*types.MsgSendTokenResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    // Call the SendToken function from the keeper to process the token transfer
    err := m.Keeper.SendToken(ctx, *msg)
    if err != nil {
        return nil, err
    }

    return &types.MsgSendTokenResponse{}, nil
}
