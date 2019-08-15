package theforcechain

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the forcechain Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Gets the entire Order metadata struct for an id
func (k Keeper) GetOrder(ctx sdk.Context, id string) Order {
	store := ctx.KVStore(k.storeKey)
	if !k.IsIdPresent(ctx, id) {
		return Order{}
	}
	bz := store.Get([]byte(id))
	var order Order
	k.cdc.MustUnmarshalBinaryBare(bz, &order)
	return order
}

// Sets the entire Order metadata struct for an id
func (k Keeper) SetOrder(ctx sdk.Context, id string, order Order) {
	if order.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(order))
}

// Deletes the entire Order metadata struct for an id
func (k Keeper) DeleteOrder(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(id))
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, id string) bool {
	return !k.GetOrder(ctx, id).Owner.Empty()
}

// // GetOwner - get the current owner of a name
// func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
// 	return k.GetWhois(ctx, name).Owner
// }

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, id string) sdk.AccAddress {
	return k.GetOrder(ctx, id).Owner
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, id string, owner sdk.AccAddress) {
	order := k.GetOrder(ctx, id)
	order.Owner = owner
	k.SetOrder(ctx, id, order)
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetIdsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// Check if the id is present in the store or not
func (k Keeper) IsIdPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(id))
}
