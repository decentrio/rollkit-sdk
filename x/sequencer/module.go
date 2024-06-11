package sequencer

import (
	"context"
	"encoding/json"
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	modulev1 "github.com/decentrio/rollkit-sdk/api/rollkitsdk/sequencer/module"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/keeper"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.AppModuleBasic  = AppModuleBasic{}
	_ appmodule.AppModule    = AppModule{}
	_ module.HasServices     = AppModule{}
	_ module.HasABCIGenesis  = AppModule{}
	_ module.HasABCIEndBlock = AppModule{}
)

type AppModuleBasic struct {
	cdc codec.Codec
}

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// Name returns the sequencer module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the sequencer module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// TODO: add validation
// ValidateGenesis performs genesis state validation for the sequencer module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the staking module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	return am.keeper.InitGenesis(ctx, &genesisState)
}

// EndBlock implements the AppModule interface
func (am AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	nextChangeSequencerHeight, err := am.keeper.NextSequencerChangeHeight.Get(ctx)
	if sdkCtx.BlockHeight() != nextChangeSequencerHeight || err != nil {
		return []abci.ValidatorUpdate{}, nil
	}

	return am.keeper.ChangeoverToRollup(sdkCtx, keeper.LastValidatorSet)
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	// am.keeper.RegisterInvariants(ir)
}
func (am AppModule) ConsensusVersion() uint64 { return 1 }

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// Name returns the sequencer module's name
func (am AppModule) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the staking module's types on the given LegacyAmino codec.
func (AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (AppModule) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the staking module.
func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// DefaultGenesis returns default genesis state as raw bytes for the staking
// module.
func (AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ExportGenesis returns the exported genesis state as raw bytes for the staking
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return nil
}

// ValidateGenesis performs genesis state validation for the staking module.
func (AppModule) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return nil
}

// start depinject implementation

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Config        *modulev1.Module
	AccountKeeper types.AccountKeeper
	Cdc           codec.Codec
	StoreService  store.KVStoreService
}

// Dependency Injection Outputs
type ModuleOutputs struct {
	depinject.Out

	SequencerKeeper keeper.Keeper
	Module          appmodule.AppModule
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
		authority.String(),
	)
	m := NewAppModule(in.Cdc, k)
	return ModuleOutputs{SequencerKeeper: k, Module: m}
}
