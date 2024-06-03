package simapp

import (
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	cmttypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	simtestutil "github.com/decentrio/rollkit-sdk/testutil/sims"
)

// SetupOptions defines arguments that are passed into `Simapp` constructor.
type SetupOptions struct {
	Logger  log.Logger
	DB      *dbm.MemDB
	AppOpts servertypes.AppOptions
}

// NewSimappWithCustomOptions initializes a new SimApp with custom options.
func NewSimappWithCustomOptions(t *testing.T, isCheckTx bool, options SetupOptions) *SimApp {
	t.Helper()
	privVal := secp256k1.GenPrivKey().PubKey()
	pubKey, err := cryptocodec.ToCmtPubKeyInterface(privVal)
	require.NoError(t, err)
	// create validator set with single validator
	validator := cmttypes.NewValidator(pubKey, 1)
	valSet := cmttypes.NewValidatorSet([]*cmttypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100000000000000))),
	}

	app := NewSimApp(options.Logger, options.DB, nil, true, options.AppOpts)
	genesisState := app.DefaultGenesis()
	genesisState, err = simtestutil.GenesisStateWithValSet(app.AppCodec(), genesisState, valSet, []authtypes.GenesisAccount{acc}, balance)
	require.NoError(t, err)

	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := cmtjson.MarshalIndent(genesisState, "", " ")
		require.NoError(t, err)

		// Initialize the chain
		_, err = app.InitChain(&abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simtestutil.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		})
		require.NoError(t, err)
	}

	return app
}
