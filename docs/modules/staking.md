---
parent: Modules
---

# Staking

The staking module in cosmos sdk is an implementation of cosmos PoS protocol. It is one of the most complex modules in the SDK which contains many logic:

- Validators management
- Delegations management
- abci valset update
- Slashing malicous validators and its delegations

With these logic, staking module plays an important role in every cosmos application, a dependency of many modules (even ibc module):

- Manage tendermint valsets (valset management and via abci.ValidatorUpdate)
- Staking protocol: delegation/redelegation/undelegation...
- Work with distribution module on allocating staking rewards (via delegations management)
- Work with government module on proposal tallying (via delegations management and valset management)
- Work with slashing module on slashing/jailing misbehaved validators and their delegations.

Rollkit is designed to use cosmos SDK for its application layer which naturally includes staking module. However, Rollkit doesn't have an actual valset or a consensus protocol which makes staking module loses its fundamental purpose. But we still shouldn't remove the staking module cause there're other modules dependant on the staking module offering useful functionalities (even for non-staking rollup) like goverment. For that reason, we decided to wrap the sdk staking module so we still keep internal functionality and apis of the module for compability, while at the same time override some parts of it to be suitable for rollkit.

This means that the staking module now doesn't implement cosmos PoS, but rather a "pseudo-staking" protocol. Thus, staking validators no longer have its previous role of making blocks, they revert to only participating in goverment (governators). Staking module will not be able to make abci valset update.

New usecases for staking module:

- Staking protocol: delegation/redelegation/undelegation...
- Work with government module on proposal tallying
- Work with distribtion module on allocating staking rewards
  
For these usecases, we make the following changes to the module:

- Remove `abci.ValidatorUpdate` from `AppModule.Endblock`.
- Remove `abci.validatorUpdate` from `AppModule.InitGenesis`.
- Remove slashing/jailing logic.

Other than that, we keep the rest of the module intact
