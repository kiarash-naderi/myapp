package keeper

import (
    "testing"
    "log"

    
    "cosmossdk.io/store"
    storetypes "cosmossdk.io/store/types"
    cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
    "github.com/cosmos/cosmos-sdk/codec"
    codectypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
    "github.com/kiarash-naderi/myapp/x/abslayer/keeper"
    myappTypes "github.com/kiarash-naderi/myapp/x/abslayer/types"
    "github.com/stretchr/testify/require"
    dbm "github.com/cosmos/cosmos-db"
    "cosmossdk.io/store/metrics"
)

type MockBankKeeper struct{}

func (mbk *MockBankKeeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
    // Mock implementation of SendCoins
    return nil
}

func AbslayerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
    storeKey := storetypes.NewKVStoreKey(myappTypes.StoreKey)

    db := dbm.NewMemDB()
    logger := log.NewNopLogger()
    storeMetrics := metrics.NewNoOpMetrics() // Use NewNoOpMetrics as a placeholder

    stateStore := store.NewCommitMultiStore(db, logger, storeMetrics)
    stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
    require.NoError(t, stateStore.LoadLatestVersion())

    registry := codectypes.NewInterfaceRegistry()
    cdc := codec.NewProtoCodec(registry)
    authority := authtypes.NewModuleAddress(govtypes.ModuleName)

    // Mocking BankKeeper for test purposes
    bankKeeper := &MockBankKeeper{}
    _ = bankKeeper // Use the variable to avoid the "declared and not used" error

    k := keeper.NewKeeper(
        cdc,
        storeKey,
        logger, // Use the adjusted logger here
        bankKeeper, // Pass the BankKeeper here
        authority.String(), // Pass the additional string argument here
    )

    ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, logger)

    // Initialize params
    if err := k.SetParams(ctx, myappTypes.DefaultParams()); err != nil {
        panic(err)
    }

    return k, ctx
}