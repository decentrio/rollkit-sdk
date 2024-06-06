package testutil

// TODO: refactor test simapp

// import (
// 	"encoding/json"
// 	"testing"
// 	"time"

// 	abci "github.com/cometbft/cometbft/abci/types"
// 	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
// 	cmttypes "github.com/cometbft/cometbft/types"
// 	"github.com/stretchr/testify/require"

// 	"cosmossdk.io/log"
// 	dbm "github.com/cosmos/cosmos-db"
// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
// 	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
// 	"github.com/cosmos/cosmos-sdk/testutil/testdata"

// 	"github.com/decentrio/rollkit-sdk/simapp"
// 	simappparams "github.com/decentrio/rollkit-sdk/simapp/params"
// 	seqtypes "github.com/decentrio/rollkit-sdk/x/sequencer/types"
// )

// const TestChainID = "testchain-1"

// var DefaultConsensusParams = &tmtypes.ConsensusParams{
// 	Block: &tmtypes.BlockParams{
// 		MaxBytes: 200000,
// 		MaxGas:   -1,
// 	},
// 	Evidence: &tmtypes.EvidenceParams{
// 		MaxAgeNumBlocks: 302400,
// 		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
// 		MaxBytes:        10000,
// 	},
// 	Validator: &tmtypes.ValidatorParams{
// 		PubKeyTypes: []string{
// 			cmttypes.ABCIPubKeyTypeEd25519,
// 		},
// 	},
// }

// // EmptyAppOptions is a stub implementing AppOptions
// type EmptyAppOptions struct{}

// // Get implements AppOptions
// func (ao EmptyAppOptions) Get(o string) interface{} {
// 	return nil
// }

// func setup(chainId string) (*simapp.SimApp, map[string]json.RawMessage) {
// 	db := dbm.NewMemDB()

// 	encCdc := simappparams.MakeTestEncodingConfig()
// 	testApp := simapp.NewSimApp(
// 		log.NewNopLogger(), db, nil, true, EmptyAppOptions{}, baseapp.SetChainID(chainId),
// 	)

// 	genState := testApp.DefaultGenesisState(encCdc.Codec)
// 	return testApp, genState
// }

// // Setup initializes a new Rollapp. A Nop logger is set in Rollapp.
// func SetupWithChainId(t *testing.T, chainId string) *simapp.SimApp {
// 	t.Helper()
// 	_, pubKey, _ := testdata.KeyTestPubAddr()

// 	pk, err := cryptocodec.ToCmtProtoPublicKey(pubKey)
// 	require.NoError(t, err)

// 	pkAny, err := codectypes.NewAnyWithValue(pubKey)
// 	require.NoError(t, err)

// 	app, genesisState := setup(chainId)

// 	// setup for sequencer
// 	seqGenesis := seqtypes.GenesisState{
// 		Params: seqtypes.Params{},
// 		Sequencers: []seqtypes.Sequencer{
// 			{
// 				Name:            "sequencer",
// 				ConsensusPubkey: pkAny,
// 			},
// 		},
// 	}
// 	genesisState[seqtypes.ModuleName] = app.AppCodec().MustMarshalJSON(&seqGenesis)

// 	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
// 	require.NoError(t, err)
// 	// init chain will set the validator set and initialize the genesis accounts
// 	_, err = app.InitChain(
// 		&abci.RequestInitChain{
// 			Time:            time.Time{},
// 			ChainId:         chainId,
// 			ConsensusParams: DefaultConsensusParams,
// 			Validators: []abci.ValidatorUpdate{
// 				{PubKey: pk, Power: 1},
// 			},
// 			AppStateBytes: stateBytes,
// 			InitialHeight: 0,
// 		},
// 	)
// 	require.NoError(t, err)

// 	return app
// }
