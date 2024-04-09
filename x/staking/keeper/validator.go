package keeper

import "github.com/cosmos/cosmos-sdk/x/staking/types"

// Implements ValidatorSet interface
var _ types.ValidatorSet = Keeper{}
