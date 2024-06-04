package keeper_test

// TODO: refactor test simapp
// func TestStore(t *testing.T) {
// 	app := testutil.SetupWithChainId(t, testutil.TestChainID)
// 	ctx := app.BaseApp.NewContext(false).WithChainID(testutil.TestChainID).WithBlockTime(time.Now().UTC()).WithBlockHeight(1)

// 	_, pubKey1, _ := testdata.KeyTestPubAddr()

// 	pkAny1, err := codectypes.NewAnyWithValue(pubKey1)
// 	require.NoError(t, err)
// 	sequencerKeeper := app.SequencerKeeper

// 	// NextSequencerChangeHeight
// 	err = sequencerKeeper.NextSequencerChangeHeight.Set(ctx, 100)
// 	require.NoError(t, err)

// 	nextSequencerChangeHeight, err := sequencerKeeper.NextSequencerChangeHeight.Get(ctx)
// 	require.NoError(t, err)
// 	require.Equal(t, nextSequencerChangeHeight, int64(100))

// 	// Sequencer
// 	err = sequencerKeeper.Sequencer.Set(ctx, types.Sequencer{
// 		Name:            "test sequence",
// 		ConsensusPubkey: pkAny1,
// 	})
// 	require.NoError(t, err)

// 	newSequencer, err := sequencerKeeper.Sequencer.Get(ctx)
// 	require.NoError(t, err)
// 	require.Equal(t, newSequencer.ConsensusPubkey, pkAny1)

// 	// Wrong type any
// 	animal := &testdata.Dog{
// 		Size_: "big",
// 		Name:  "Dog",
// 	}
// 	anyAnimal, err := codectypes.NewAnyWithValue(animal)
// 	require.NoError(t, err)

// 	err = sequencerKeeper.Sequencer.Set(ctx, types.Sequencer{
// 		Name:            "test sequence",
// 		ConsensusPubkey: anyAnimal,
// 	})
// 	require.NoError(t, err)

// 	_, err = sequencerKeeper.Sequencer.Get(ctx)
// 	require.Error(t, err)
// }
