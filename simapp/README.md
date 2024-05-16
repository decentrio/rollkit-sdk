# Default app template

Here we have the template for a vanilla rollkit application using our rappdk. From this template, rollkit developers can begin to customize their own application by importing modules into the app.

The reasonale behind this is that the typical cosmos-sdk app is designed for cosmos-PoS, meaning it has many redundant features for rollkit. Keeping these features makes the rollkit application confusing and unnecessary complicated.

List of modules from the cosmos-sdk that should be removed for a rollkit application:

- Slashing module
    This module serves as a slashing mechanism against malicious/misbehaving actors in PoS consensus protocol, not needed for rollkit.
- Evidence module
    This module receives evidences of malicious/misbehaving actors in PoS consensus protocol and trigger slashing, not needed for rollkit.
- Mint module
    This module is responsible for generating block reward to incentivize PoS validators, not needed for rollkit.

Besides the removal of some sdk modules, we also make the following changes:


add a module `x/sequencer` to manage