# Integrate the RappDK to your chain

This guide will teach you how to integrate RappDK to your cometBFT chain.

## Go mod

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

## Configuring and Adding wrapper staking

1. Add the following modules to `app.go`

    ```go
    import (
    ... 
        rollkitstaking "github.com/decentrio/rollkit-sdk/x/staking"
        rollkitstakingkeeper "github.com/decentrio/rollkit-sdk/x/staking/keeper"
    ...
    )
    ```

2. Replace staking AppModule by RappDK staking
    In `app.ModuleManager` initial
    replace

    ```go
    staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),

    ```

    by

    ```go
    rollkitstaking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),

    ```

3. Replace Cosmos-SDK staking keeper by RappDK staking keeper
    In `app.StakingKeeper` initial
    replace

    ```go
    app.StakingKeeper = stakingkeeper.NewKeeper(

  appCodec, runtime.NewKVStoreService(keys[stakingtypes.StoreKey]), app.AccountKeeper, app.BankKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(), authcodec.NewBech32Codec(sdk.Bech32PrefixValAddr), authcodec.NewBech32Codec(sdk.Bech32PrefixConsAddr))

    ```

    by

    ```go
    app.StakingKeeper = rollkitstakingkeeper.NewKeeper(
  appCodec, runtime.NewKVStoreService(keys[stakingtypes.StoreKey]), app.AccountKeeper, app.BankKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(), authcodec.NewBech32Codec(sdk.Bech32PrefixValAddr), authcodec.NewBech32Codec(sdk.Bech32PrefixConsAddr))

    ```

## Configuring and adding sequencer module

In this step, you should add sequencer module to your `app.go` like other normal module in cosmos-SDK.
