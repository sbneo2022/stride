package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Stride-Labs/stride/v16/x/stakeibc/types"
)

// SetTradeRoute set a specific tradeRoute in the store
func (k Keeper) SetTradeRoute(ctx sdk.Context, tradeRoute types.TradeRoute) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeRouteKeyPrefix))
	key := tradeRoute.GetKey()
	b := k.cdc.MustMarshal(&tradeRoute)
	store.Set(key, b)
}

// GetTradeRoute returns a tradeRoute from its start and end denoms
func (k Keeper) GetTradeRoute(ctx sdk.Context, startDenom string, endDenom string) (val types.TradeRoute, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeRouteKeyPrefix))
	key := types.TradeRouteKeyFromDenoms(startDenom, endDenom)
	b := store.Get(key)
	if len(b) == 0 {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradeRoute removes a tradeRoute from the store
func (k Keeper) RemoveTradeRoute(ctx sdk.Context, startDenom string, endDenom string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeRouteKeyPrefix))
	key := types.TradeRouteKeyFromDenoms(startDenom, endDenom)
	store.Delete(key)
}

// GetAllTradeRoute returns all tradeRoutes
func (k Keeper) GetAllTradeRoutes(ctx sdk.Context) (list []types.TradeRoute) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeRouteKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradeRoute
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}