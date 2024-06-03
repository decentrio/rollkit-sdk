package upgrade

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sequencertypes "github.com/decentrio/rollkit-sdk/x/sequencer/types"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	sequencerkeeper "github.com/decentrio/rollkit-sdk/x/sequencer/keeper"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
)

// Name is upgrade name.
const Name = "rollup-migrate"

var StoreUpgrades = storetypes.StoreUpgrades{
	Added: []string{
		sequencertypes.ModuleName,
	},
	Deleted: []string{},
}

func CreateUpgradeHandler(mm *module.Manager, configurator module.Configurator, seqKeeper sequencerkeeper.Keeper, sk stakingkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		seqPubkey := "J3ZVpvQv18IveVGkRuW+Yog9R/7E4gTWLzWIRiOw9Zk="

		sdkCtx := sdk.UnwrapSDKContext(ctx)
		// get last validator set
		validatorSet, err := sk.GetLastValidators(ctx)
		if err != nil {
			return nil, err
		}

		pubKey, err := GetSequencerEd25519Pubkey(seqPubkey)
		if err != nil {
			return nil, err
		}

		pkAny, err := codectypes.NewAnyWithValue(pubKey)
		if err != nil {
			return nil, err
		}
		err = seqKeeper.Sequencer.Set(sdkCtx, types.Sequencer{
			Name:            "sequencer",
			ConsensusPubkey: pkAny,
		})
		if err != nil {
			return nil, err
		}

		sequencerkeeper.LastValidatorSet = validatorSet
		err = seqKeeper.NextSequencerChangeHeight.Set(sdkCtx, sdkCtx.BlockHeight())
		if err != nil {
			return nil, err
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
