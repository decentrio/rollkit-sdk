package upgrade

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sequencerkeeper "github.com/decentrio/rollkit-sdk/x/sequencer/keeper"
)

// Name is migration name.
const Name = "rollup-migrate"

func CreateUpgradeHandler(mm *module.Manager, configurator module.Configurator, seqKeeper sequencerkeeper.Keeper, sk stakingkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		// get last validator set
		validatorSet, err := sk.GetLastValidators(ctx)
		if err != nil {
			return nil, err
		}

		sequencerkeeper.LastValidatorSet = validatorSet
		seqKeeper.SetNextSequencerChangeHeight(sdkCtx, sdkCtx.BlockHeight())

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
