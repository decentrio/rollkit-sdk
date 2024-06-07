# Default app template

Here we have the [template](app.go) for a vanilla rollkit application using our rappdk. From this template, rollkit developers can begin to customize their own application by importing modules into the app.

The reasonale behind this is that the typical cosmos-sdk app is designed for cosmos-PoS, meaning it has many redundant features for rollkit. Keeping these features makes the rollkit application confusing and unnecessary complicated.

List of modules from the cosmos-sdk that should be removed for a rollkit application:

- Slashing module
    This module serves as a slashing mechanism against malicious/misbehaving actors in PoS consensus protocol, not needed for rollkit.
- Evidence module
    This module receives evidences of malicious/misbehaving actors in PoS consensus protocol and trigger slashing, not needed for rollkit.
- Mint module
    This module is responsible for generating block reward to incentivize PoS validators, not needed for rollkit.

Besides the removal of some sdk modules, we also make the following changes to the app:

- Wrap `x/staking` module to modify it without breaking related modules. The goal is to simplify the module, transforming it to a `pseudo-staking` module that can still incoperate with other sdk modules. Read this [doc](./modules/staking.md) for more details

- Add module `x/sequencer` to manage the sequencer as `x/staking` no longer has the role of managing the underlying valset(equivalent to the sequencer). Read this [doc](./modules/sequencer.md) for more details

The default app is designed to be a very simple rollkit application without any flavors. It has the very same functionalities as a cosmos-sdk app but with PoS related features removed (slashing and PoS incentives). Though, it still has a `staking` app with empty incentive settings because we want to provide the base for staking rollkit apps to build ontop. Even if the app doesn't want `staking` feature, the `staking` app still serves as the base for gov logic which is still very useful.
