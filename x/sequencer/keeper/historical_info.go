package keeper

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

// GetHistoricalInfo gets the historical info at a given height
func (k Keeper) GetHistoricalInfo(ctx context.Context, height int64) (types.HistoricalInfo, error) {
	store := k.storeService.OpenKVStore(ctx)
	key := types.GetHistoricalInfoKey(height)

	value, err := store.Get(key)
	if err != nil {
		return types.HistoricalInfo{}, err
	}

	if value == nil {
		return types.HistoricalInfo{}, types.ErrNoHistoricalInfo
	}

	return types.UnmarshalHistoricalInfo(k.cdc, value)
}

// SetHistoricalInfo sets the historical info at a given height
func (k Keeper) SetHistoricalInfo(ctx context.Context, height int64, hi *types.HistoricalInfo) error {
	store := k.storeService.OpenKVStore(ctx)
	key := types.GetHistoricalInfoKey(height)
	value, err := k.cdc.Marshal(hi)
	if err != nil {
		return err
	}
	return store.Set(key, value)
}

// DeleteHistoricalInfo deletes the historical info at a given height
func (k Keeper) DeleteHistoricalInfo(ctx context.Context, height int64) error {
	store := k.storeService.OpenKVStore(ctx)
	key := types.GetHistoricalInfoKey(height)

	return store.Delete(key)
}

// IterateHistoricalInfo provides an iterator over all stored HistoricalInfo
// objects. For each HistoricalInfo object, cb will be called. If the cb returns
// true, the iterator will break and close.
func (k Keeper) IterateHistoricalInfo(ctx context.Context, cb func(types.HistoricalInfo) bool) error {
	store := k.storeService.OpenKVStore(ctx)
	iterator, err := store.Iterator(types.HistoricalInfoKey, storetypes.PrefixEndBytes(types.HistoricalInfoKey))
	if err != nil {
		return err
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		histInfo, err := types.UnmarshalHistoricalInfo(k.cdc, iterator.Value())
		if err != nil {
			return err
		}
		if cb(histInfo) {
			break
		}
	}

	return nil
}

// TrackHistoricalInfo saves the latest historical-info and deletes the oldest
// heights that are below pruning height
func (k Keeper) TrackHistoricalInfo(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	historicalEntry := types.HistoricalInfo{
		Header: sdkCtx.BlockHeader(),
	}

	// Set latest HistoricalInfo at current height
	return k.SetHistoricalInfo(ctx, sdkCtx.BlockHeight(), &historicalEntry)
}
