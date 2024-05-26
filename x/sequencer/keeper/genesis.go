package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"

	abci "github.com/cometbft/cometbft/abci/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) []abci.ValidatorUpdate {
	var seq types.Sequencer
	if len(data.Sequencers) == 0 {
		return []abci.ValidatorUpdate{}
	}

	// Set the initial sequence
	k.SetSequencer(ctx, data.Sequencers[0])
	seq = k.GetSequencer(ctx)

	if seq == (types.Sequencer{}) {
		panic("Sequencer not set")
	}

	pk, err := seq.TmConsPublicKey()
	if err != nil {
		panic(err)
	}

	return []abci.ValidatorUpdate{
		{
			PubKey: pk,
			Power:  1,
		},
	}
}
