package keeper

import (
	"testing"

	"cosmossdk.io/log"            // Import Cosmos SDK logger
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/kiarash-naderi/myapp/x/abslayer/keeper"
	myappTypes "github.com/kiarash-naderi/myapp/x/abslayer/types"
	"github.com/stretchr/testify/require"
)

type MockBankKeeper struct{}

func (mbk *MockBankKeeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	// Mock implementation of SendCoins
	return nil
}

func AbslayerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(myappTypes.StoreKey)

	db := dbm.NewMemDB()

	// Use Cosmos SDK logger directly
	logger := log.NewNopLogger() // Cosmos SDK logger that doesn't output anything

	storeMetrics := metrics.NewNoOpMetrics() // Initialize storeMetrics
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
		logger,           // Use Cosmos SDK logger
		bankKeeper,       // Pass the BankKeeper here
		authority.String(), // Pass the additional string argument here
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, logger)

	// Initialize params
	if err := k.SetParams(ctx, myappTypes.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}
