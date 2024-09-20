package abslayer

import (
    "context"
    "encoding/json"
    "fmt"

    "cosmossdk.io/core/appmodule"
    "cosmossdk.io/core/store"
    "cosmossdk.io/depinject"
    "cosmossdk.io/log"
    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    sdkstore "cosmossdk.io/store/types"

    modulev1 "github.com/kiarash-naderi/myapp/api/myapp/abslayer/module"
    "github.com/kiarash-naderi/myapp/x/abslayer/keeper"
    "github.com/kiarash-naderi/myapp/x/abslayer/types"

    
	
)

var (
    _ module.AppModuleBasic      = (*AppModule)(nil)
    _ module.AppModuleSimulation = (*AppModule)(nil)
    _ module.HasGenesis          = (*AppModule)(nil)
    _ module.HasInvariants       = (*AppModule)(nil)
    _ module.HasConsensusVersion = (*AppModule)(nil)

    _ appmodule.AppModule       = (*AppModule)(nil)
    _ appmodule.HasBeginBlocker = (*AppModule)(nil)
    _ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

type AppModuleBasic struct {
    cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
    return AppModuleBasic{cdc: cdc}
}

func (AppModuleBasic) Name() string {
    return types.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
    types.RegisterInterfaces(reg)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    return cdc.MustMarshalJSON(types.DefaultGenesis())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
    var genState types.GenesisState
    if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
        return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
    }
    return genState.Validate()
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
    if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
        panic(err)
    }
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

type AppModule struct {
    AppModuleBasic

    keeper        keeper.Keeper
    accountKeeper types.AccountKeeper
    bankKeeper    types.BankKeeper
}

func NewAppModule(
    cdc codec.Codec,
    keeper keeper.Keeper,
    accountKeeper types.AccountKeeper,
    bankKeeper types.BankKeeper,
) AppModule {
    return AppModule{
        AppModuleBasic: NewAppModuleBasic(cdc),
        keeper:         keeper,
        accountKeeper:  accountKeeper,
        bankKeeper:     bankKeeper,
    }
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
    types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
    types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
    var genState types.GenesisState
    cdc.MustUnmarshalJSON(gs, &genState)
    InitGenesis(ctx, am.keeper, genState)
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
    genState := ExportGenesis(ctx, am.keeper)
    return cdc.MustMarshalJSON(genState)
}

func (AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) BeginBlock(ctx context.Context) error {
    // Implement your begin block logic here
    return nil
}

func (am AppModule) EndBlock(ctx context.Context) error {
    // Implement your end block logic here
    return nil
}

func (am AppModule) IsOnePerModuleType() {}

func (am AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

func init() {
    appmodule.Register(
        &modulev1.Module{},
        appmodule.Provide(ProvideModule),
    )
}

type ModuleInputs struct {
    depinject.In

    StoreService store.KVStoreService
    Cdc          codec.Codec
    Config       *modulev1.Module
    Logger       log.Logger

    AccountKeeper types.AccountKeeper
    BankKeeper    types.BankKeeper
}

type ModuleOutputs struct {
    depinject.Out

    AbslayerKeeper keeper.Keeper
    Module         appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
    authority := authtypes.NewModuleAddress(govtypes.ModuleName)
    if in.Config.Authority != "" {
        authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
    }

    storeKey := sdkstore.NewKVStoreKey("my_store_key")
    k := keeper.NewKeeper(
        in.Cdc,
        storeKey,
        in.Logger.With("module", "abslayer"),
        in.BankKeeper,
        authority.String(),
    )

    m := NewAppModule(
        in.Cdc,
        k,
        in.AccountKeeper,
        in.BankKeeper,
    )

    in.Logger.Info("Finished providing AppModule for Abslayer module")
    return ModuleOutputs{AbslayerKeeper: k, Module: m}
}