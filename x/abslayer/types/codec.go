package types

import (
    "github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/msgservice"
)

// ModuleCdc defines the codec to be used by the module
var (
    ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// RegisterInterfaces registers the necessary interfaces for the message
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    registry.RegisterImplementations(
        (*sdk.Msg)(nil),
        &MsgSendToken{},
    )
    msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
