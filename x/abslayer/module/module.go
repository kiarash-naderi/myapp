package abslayer

import (


"testing"

    cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
    "github.com/cosmos/cosmos-sdk/codec"
    codectypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/kiarash-naderi/myapp/x/abslayer/keeper"
    myappTypes "github.com/kiarash-naderi/myapp/x/abslayer/types"
    "github.com/stretchr/testify/require"
    dbm "github.com/cosmos/cosmos-db"
    "cosmossdk.io/log" 
    
    storetypes "cosmossdk.io/store/types" // Import the storetypes package from cosmossdk.io
    "cosmossdk.io/store" // Import the store package from cosmossdk.io
    "cosmossdk.io/store/metrics" // Import the metrics package from cosmossdk.io
)

type MockBankKeeper struct{}

func (mbk *MockBankKeeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
    // Mock implementation of SendCoins
    return nil
}

func AbslayerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
    storeKey := storetypes.NewKVStoreKey(myappTypes.StoreKey)

    db := dbm.NewMemDB()
   
    storeMetrics := metrics.NewNoOpMetrics() // Use NewNoOpMetrics as a placeholder

    logger := log.NewNopLogger()  // Use the standard log package with required arguments

    stateStore := store.NewCommitMultiStore(db, logger, storeMetrics)
    stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
    require.NoError(t, stateStore.LoadLatestVersion())

    registry := codectypes.NewInterfaceRegistry()
    cdc := codec.NewProtoCodec(registry)
    // authority := authtypes.NewModuleAddress(govtypes.ModuleName)

    // Mocking BankKeeper for test purposes
    bankKeeper := &MockBankKeeper{}
    _ = bankKeeper // Use the variable to avoid the "declared and not used" error

    // Create the keeper using the codec
    stdLogger := log.NewNopLogger() // Initialize the logger
    k := keeper.NewKeeper(cdc, storeKey, stdLogger, bankKeeper, myappTypes.ModuleName)

   
  
    ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, logger)

    // Initialize params
    if err := k.SetParams(ctx, myappTypes.DefaultParams()); err != nil {
        panic(err)
    }

    return k, ctx
}