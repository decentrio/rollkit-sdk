package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
)

// TODO: testing
func (keeper Keeper) SetSequencer(ctx sdk.Context, sequencer types.Sequencer) {
	store := keeper.storeService.OpenKVStore(ctx)
	store.Set(types.SequencerConsAddrKey, keeper.cdc.MustMarshal(&sequencer))
}

// TODO: testing
func (keeper Keeper) GetSequencer(ctx sdk.Context) (seq types.Sequencer) {
	store := keeper.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.SequencerConsAddrKey)
	if err != nil {
		return types.Sequencer{}
	}
	keeper.cdc.MustUnmarshal(bz, &seq)
	return seq
}

// TODO: testing
func (keeper Keeper) SetNextSequencerChangeHeight(ctx sdk.Context, height int64) {
	store := keeper.storeService.OpenKVStore(ctx)

	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, uint64(height))

	store.Set(types.SequencerConsAddrKey, bz)
}

// TODO: testing
func (keeper Keeper) GetNextSequencerChangeHeight(ctx sdk.Context) (int64, error) {
	store := keeper.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.NextSequencerChangeHeight)
	if err != nil {
		return 0, fmt.Errorf("no plan for changing sequencer")
	}

	return int64(binary.LittleEndian.Uint64(bz)), nil
}

// TODO: testing
func (keeper Keeper) RemoveNextSequencerChangeHeight(ctx sdk.Context) error {
	store := keeper.storeService.OpenKVStore(ctx)
	return store.Delete(types.NextSequencerChangeHeight)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
