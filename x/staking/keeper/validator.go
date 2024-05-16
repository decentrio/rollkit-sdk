package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Implements ValidatorSet interface
var _ types.ValidatorSet = Keeper{}

// Slash a validator for an infraction committed at a known height.
// Since this module implements pseudo-staking, Slash performs no-op
func (k Keeper) Slash(context.Context, sdk.ConsAddress, int64, int64, math.LegacyDec) (math.Int, error) {
	return math.ZeroInt(), nil
}

// SlashWithInfractionReason performs no-op
func (k Keeper) SlashWithInfractionReason(context.Context, sdk.ConsAddress, int64, int64, math.LegacyDec, types.Infraction) (math.Int, error) {
	return math.ZeroInt(), nil
}

// Jail performs no-op
func (k Keeper) Jail(context.Context, sdk.ConsAddress) error {
	return nil
}

// Unjail performs no-op
func (k Keeper) Unjail(context.Context, sdk.ConsAddress) error {
	return nil
}
