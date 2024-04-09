package keeper

import (
	"context"
	"errors"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// disable delegate and create new validator
func (k Keeper) Delegate(
	ctx context.Context, delAddr sdk.AccAddress, bondAmt math.Int, tokenSrc stakingtypes.BondStatus,
	validator stakingtypes.Validator, subtractAccount bool,
) (newShares math.LegacyDec, err error) {
	return math.LegacyDec{}, errors.New("not implemented")
}

// disable redelegate
func (k Keeper) BeginRedelegation(
	ctx context.Context, delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress, sharesAmount math.LegacyDec,
) (completionTime time.Time, err error) {
	return time.Time{}, errors.New("not implemented")
}

// disable undelegate
func (k Keeper) Undelegate(
	ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, sharesAmount math.LegacyDec,
) (time.Time, math.Int, error) {
	return time.Time{}, math.Int{}, errors.New("not implemented")
}
