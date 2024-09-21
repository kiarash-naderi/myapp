package keeper

import (
    "errors"
    "fmt"
    
    "cosmossdk.io/log"
    storetypes "cosmossdk.io/store/types"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"



    "github.com/kiarash-naderi/myapp/x/abslayer/types"
)

type (
    Keeper struct {
        cdc        codec.BinaryCodec
        storeKey   storetypes.StoreKey
        logger     log.Logger   
        bankKeeper types.BankKeeper
        authority  string
    }
)


func NewKeeper(
    cdc codec.BinaryCodec,
    storeKey storetypes.StoreKey,
    logger log.Logger,  // logger SDK ro estefade kon
    bankKeeper types.BankKeeper,
    authority string,
) Keeper {
    if _, err := sdk.AccAddressFromBech32(authority); err != nil {
        panic(fmt.Sprintf("invalid authority address: %s", authority))
    }

    return Keeper{
        cdc:        cdc,
        storeKey:   storeKey,
        logger:     logger,   
        bankKeeper: bankKeeper,
        authority:  authority,
    }
}


// Data storage methods
func (k *Keeper) StoreTransaction(ctx sdk.Context, tx types.Transaction) {
    store := ctx.KVStore(k.storeKey)
    bz := k.cdc.MustMarshal(&tx)
    store.Set([]byte(tx.ID), bz)
}

func (k Keeper) GetAuthority() string {
    return k.authority
}

func (k Keeper) GetTransaction(ctx sdk.Context, id string) (types.Transaction, error) {
    store := ctx.KVStore(k.storeKey)
    bz := store.Get([]byte(id))
    if bz == nil {
        return types.Transaction{}, errors.New("transaction not found")
    }
    var tx types.Transaction
    k.cdc.MustUnmarshal(bz, &tx)
    return tx, nil
}