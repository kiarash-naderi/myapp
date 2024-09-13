package keeper

import (
	"errors"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kiarash-naderi/myapp/x/abslayer/types"
)

type (
    Keeper struct {
        cdc          codec.BinaryCodec
        storeService store.KVStoreService
        logger       log.Logger
        bankKeeper   types.BankKeeper // This connects to the bank module to send coins
        authority    string
    }
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeService store.KVStoreService,
    logger log.Logger,
    bankKeeper types.BankKeeper, // Add the bank keeper here
    authority string,

) Keeper {
    if _, err := sdk.AccAddressFromBech32(authority); err != nil {
        panic(fmt.Sprintf("invalid authority address: %s", authority))
    }

    return Keeper{
        cdc:          cdc,
        storeService: storeService,
        logger:       logger,
        bankKeeper:   bankKeeper, // Initialize the bank keeper
        authority:    authority,
    }
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
    return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
    return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SendToken(ctx sdk.Context, msg types.MsgSendToken) error {
    sender, err := sdk.AccAddressFromBech32(msg.Sender)
    if err != nil {
        return errors.New("invalid sender address")
    }

    receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
    if err != nil {
        return errors.New("invalid receiver address")
    }

    // Send tokens
    err = k.bankKeeper.SendCoins(ctx, sender, receiver, msg.Amount)
    if err != nil {
        return err
    }

    return nil
}