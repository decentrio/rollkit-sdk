package keeper_test

// func NewTestInterfaceRegistry() codectypes.InterfaceRegistry {
// 	registry := codectypes.NewInterfaceRegistry()
// 	registry.RegisterInterface("Animal", (*testdata.Animal)(nil))
// 	registry.RegisterImplementations(
// 		(*testdata.Animal)(nil),
// 		&testdata.Dog{},
// 		&testdata.Cat{},
// 	)
// 	return registry
// }

// func TestSequencerInitGenesis(t *testing.T) {
// 	app := testutil.SetupWithChainId(t, testutil.TestChainID)
// 	ctx := app.BaseApp.NewContext(false).WithChainID(testutil.TestChainID).WithBlockTime(time.Now().UTC()).WithBlockHeight(1)

// 	_, pubKey1, _ := testdata.KeyTestPubAddr()
// 	_, pubKey2, _ := testdata.KeyTestPubAddr()

// 	pk1, err := cryptocodec.ToCmtProtoPublicKey(pubKey1)
// 	require.NoError(t, err)

// 	pkAny1, err := codectypes.NewAnyWithValue(pubKey1)
// 	require.NoError(t, err)
// 	pkAny2, err := codectypes.NewAnyWithValue(pubKey2)
// 	require.NoError(t, err)

// 	sequencerKeeper := app.SequencerKeeper
// 	genesisState := seqtypes.GenesisState{}

// 	// Empty sequencers
// 	valsUpdate := sequencerKeeper.InitGenesis(ctx, &genesisState)
// 	require.Len(t, valsUpdate, 0)

// 	// Wrong type any
// 	animal := &testdata.Dog{
// 		Size_: "big",
// 		Name:  "Dog",
// 	}
// 	anyAnimal, err := codectypes.NewAnyWithValue(animal)
// 	require.NoError(t, err)

// 	genesisState.Sequencers = []seqtypes.Sequencer{
// 		{
// 			Name:            "malicious sequencer",
// 			ConsensusPubkey: anyAnimal,
// 		},
// 	}

// 	t.Run("test init genesis with wrong", func(t *testing.T) {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				// This is due to the any types Dog does not register interfaces PubAny
// 				// in register registry, then the get sequencer return empty sequencer
// 				require.Equal(t, r, "Sequencer not set")
// 			}
// 		}()
// 		_ = sequencerKeeper.InitGenesis(ctx, &genesisState)
// 	})

// 	// With sequencer
// 	genesisState.Sequencers = []seqtypes.Sequencer{
// 		{
// 			Name:            "sequencer1",
// 			ConsensusPubkey: pkAny1,
// 		},
// 		{
// 			Name:            "sequencer2",
// 			ConsensusPubkey: pkAny2,
// 		},
// 	}
// 	valsUpdate = sequencerKeeper.InitGenesis(ctx, &genesisState)
// 	require.Len(t, valsUpdate, 1)
// 	require.Equal(t, valsUpdate[0].Power, int64(1))
// 	require.Equal(t, valsUpdate[0].PubKey, pk1)
// }
