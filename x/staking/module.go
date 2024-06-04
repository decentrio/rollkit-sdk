package staking

import (
	"context"
	"encoding/json"

	modulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"
	"github.com/decentrio/rollkit-sdk/x/staking/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
	_ module.HasServices         = AppModule{}
	_ module.HasInvariants       = AppModule{}
	_ module.HasABCIEndBlock     = AppModule{}
	_ module.HasGenesis          = AppModule{}

	_ appmodule.AppModule       = AppModule{}
	_ appmodule.HasBeginBlocker = AppModule{}
)

type AppModuleBasic struct {
	staking.AppModuleBasic
}

// AppModule embeds the Cosmos SDK's x/staking AppModule where we only override specific methods.
type AppModule struct {
	staking.AppModule
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper, ls exported.Subspace) AppModule {
	return AppModule{
		AppModule: staking.NewAppModule(cdc, &keeper.Keeper, ak, bk, ls),
		keeper:    keeper,
	}
}

func (am AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	_, err := am.keeper.EndBlocker(ctx)
	if err != nil {
		return nil, err
	}

	return []abci.ValidatorUpdate{}, nil
}

// BeginBlock returns the begin blocker for the staking module.
func (am AppModule) BeginBlock(ctx context.Context) error {
	return am.keeper.BeginBlocker(ctx)
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	_ = am.keeper.InitGenesis(ctx, &genesisState)
}

// start depinject implementation
type ModuleInputs struct {
	depinject.In

	Config                *modulev1.Module
	ValidatorAddressCodec runtime.ValidatorAddressCodec
	ConsensusAddressCodec runtime.ConsensusAddressCodec
	AccountKeeper         types.AccountKeeper
	BankKeeper            types.BankKeeper
	Cdc                   codec.Codec
	StoreService          store.KVStoreService

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace exported.Subspace `optional:"true"`
}

// Dependency Injection Outputs
type ModuleOutputs struct {
	depinject.Out

	StakingKeeper *keeper.Keeper
	Module        appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.AccountKeeper,
		in.BankKeeper,
		authority.String(),
		in.ValidatorAddressCodec,
		in.ConsensusAddressCodec,
	)
	m := NewAppModule(in.Cdc, k, in.AccountKeeper, in.BankKeeper, in.LegacySubspace)
	return ModuleOutputs{StakingKeeper: &k, Module: m}
}
