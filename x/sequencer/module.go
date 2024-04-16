package sequencer

import (
	"context"
	"encoding/json"

	"cosmossdk.io/core/appmodule"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/decentrio/rollkit-sdk/x/sequencer/keeper"
	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ appmodule.AppModule   = AppModule{}
)

type AppModuleBasic struct {
	staking.AppModuleBasic
}

type AppModule struct {
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		keeper: keeper,
	}
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	return am.keeper.InitGenesis(ctx, genesisState)
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	// am.keeper.RegisterInvariants(ir)
}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// Name returns the sequencer module's name.
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

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the staking module.
func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}
