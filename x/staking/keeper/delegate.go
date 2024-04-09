package keeper

import "github.com/cosmos/cosmos-sdk/x/staking/types"

// Implements DelegationSet interface
var _ types.DelegationSet = Keeper{}
