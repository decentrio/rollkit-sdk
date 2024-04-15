package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
)

func (keeper Keeper) SetSequencer(ctx sdk.Context, sequencer types.Sequencer) {
	store := keeper.storeService.OpenKVStore(ctx)
	store.Set(types.SequencerConsAddrKey, keeper.cdc.MustMarshal(&sequencer))
}

func (keeper Keeper) GetSequencer(ctx sdk.Context) (seq types.Sequencer) {
	store := keeper.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.SequencerConsAddrKey)
	if err != nil {
		return types.Sequencer{}
	}

	keeper.cdc.MustUnmarshal(bz, &seq)
	return seq
}
