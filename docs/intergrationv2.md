# Integrate the RappDK to your chain

This guide will teach you how to integrate RappDK to your cometBFT chain using depinject.


1. Add the rollkit-sdk package to the go.mod and install it.

    ```go
    require (
    ...
    github.com/decentrio/rollkit-sdk v<VERSION>
    ...
    )
    ```

    Use our fork rollkit version:


    ```go
    github.com/rollkit/rollkit => github.com/decentrio/rollkit v0.0.0-20240516071120-d40857416a55s
    ```

**Notice: Migration requires Rollkit to allow ABCI valset changes so using our fork version is for this. We're working with Rollkit team for upstream this feature ! [Issue Link](https://github.com/rollkit/rollkit/issues/1673) !**

2. Add sequencer and staking module into appconfig.go

In this step, you should add sequencer and wrapper staking  module to your `app.go` and `appconfig.go` like other normal module in cosmos-SDK.

We have instruction here: https://docs.cosmos.network/main/build/building-apps/app-go-v2.

3. Add upgrade handler

Example upgrade handler: 
```
func CreateUpgradeHandler(mm *module.Manager, configurator module.Configurator, seqKeeper sequencerkeeper.Keeper, sk stakingkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		seqPubkey := "J3ZVpvQv18IveVGkRuW+Yog9R/7E4gTWLzWIRiOw9Zk="

		sdkCtx := sdk.UnwrapSDKContext(ctx)
		// get last validator set
		validatorSet, err := sk.GetLastValidators(ctx)
		if err != nil {
			return nil, err
		}

		pubKey, err := GetSequencerEd25519Pubkey(seqPubkey)
		if err != nil {
			return nil, err
		}

		pkAny, err := codectypes.NewAnyWithValue(pubKey)
		if err != nil {
			return nil, err
		}
		err = seqKeeper.Sequencer.Set(sdkCtx, types.Sequencer{
			Name:            "sequencer",
			ConsensusPubkey: pkAny,
		})
		if err != nil {
			return nil, err
		}

		sequencerkeeper.LastValidatorSet = validatorSet
		err = seqKeeper.NextSequencerChangeHeight.Set(sdkCtx, sdkCtx.BlockHeight())
		if err != nil {
			return nil, err
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

```