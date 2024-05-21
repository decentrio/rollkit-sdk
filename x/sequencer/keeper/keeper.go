package keeper

import (
	"cosmossdk.io/log"
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
)

type Keeper struct {
	storeService storetypes.KVStoreService
	cdc          codec.BinaryCodec

	authKeeper types.AccountKeeper
	authority  string

	Schema                    collections.Schema
	Sequencer                 collections.Item[types.Sequencer]
	NextSequencerChangeHeight collections.Item[int64]
	Params                    collections.Item[types.Params]
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

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		storeService:              storeService,
		cdc:                       cdc,
		authKeeper:                ak,
		authority:                 authority,
		Params:                    collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Sequencer:                 collections.NewItem(sb, types.SequencerConsAddrKey, "sequencer", codec.CollValue[types.Sequencer](cdc)),
		NextSequencerChangeHeight: collections.NewItem(sb, types.NextSequencerChangeHeight, "next_sequencer_change_height", collections.Int64Value),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (keeper Keeper) GetSequencer(ctx sdk.Context) types.Sequencer {
	seq, err := keeper.Sequencer.Get(ctx)
	if err != nil {
		return types.Sequencer{}
	}
	return seq
}