package keeper

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
)

var LastValidatorSet []stakingtypes.Validator

func (k Keeper) MigrateFromSoveregin(ctx sdk.Context, sequencer types.Sequencer) error {
	// Migrate state from sovereign chain
	return k.Sequencer.Set(ctx, sequencer)
}

// ChangeoverToConsumer includes the logic that needs to execute during the process of a
// cometBFT chain to rollup changeover. This method constructs validator updates
// that will be given to tendermint, which allows the consumer chain to
// start using the provider valset, while the standalone valset is given zero voting power where appropriate.
func (k Keeper) ChangeoverToRollup(ctx sdk.Context, lastValidatorSet []stakingtypes.Validator) (initialValUpdates []abci.ValidatorUpdate, err error) {
	seq := k.GetSequencer(ctx)
	pk, err := seq.TmConsPublicKey()
	if err != nil {
		return nil, err
	}
	sequencersUpdate := abci.ValidatorUpdate{
		PubKey: pk,
		Power:  1,
	}

	for _, val := range lastValidatorSet {
		powerUpdate := val.ABCIValidatorUpdateZero()
		if val.ConsensusPubkey.Equal(seq.ConsensusPubkey) {
			continue
		}
		initialValUpdates = append(initialValUpdates, powerUpdate)
	}

	LastValidatorSet = nil
	k.Logger(ctx).Info("Rollup changeover complete - you are now a rollup chain!")

	err = k.NextSequencerChangeHeight.Remove(ctx)
	if err != nil {
		return nil, err
	}

	return append(initialValUpdates, sequencersUpdate), nil
}
