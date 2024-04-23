package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"

	abci "github.com/cometbft/cometbft/abci/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) []abci.ValidatorUpdate {
	if len(data.Sequencers) != 1 {
		panic("Genesis state must contain exactly one sequencer")
	}
	// Set the initial sequence
	k.SetSequencer(ctx, data.Sequencers[0])
	pk, err := data.Sequencers[0].TmConsPublicKey()
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
