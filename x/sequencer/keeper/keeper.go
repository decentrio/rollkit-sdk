package keeper

import (
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
)

type Keeper struct {
	storeService storetypes.KVStoreService
	cdc          codec.BinaryCodec

	authKeeper types.AccountKeeper
	authority  string
}

// NewKeeper creates a new sequencer Keeper instance
func NewKeeper(cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	ak types.AccountKeeper,
	authority string,
) Keeper {
	// ensure that authority is a valid AccAddress
	if _, err := ak.AddressCodec().StringToBytes(authority); err != nil {
		panic("authority is not a valid acc address")
	}

	return Keeper{
		storeService: storeService,
		cdc:          cdc,
		authKeeper:   ak,
		authority:    authority,
	}
}
