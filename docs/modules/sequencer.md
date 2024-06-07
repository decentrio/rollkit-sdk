---
parent: Modules
---

# Sequencer

Since the x/staking now only manages the governators, it should no longer be meant to manage the actual valset. This means we should have another module to handle said task (separation of concerns) so that we still keep abci semantics intact (enable abci valset update for rollkit).

Sequencer is the module that manages the actual valset (the sequencer). It's the only module that can make `abci valset update` to rollkit. When initializing a rollkit rollup, we now have the sequencer module init the actual valset (sequencer) instead of the staking module.

This module can potentially be integrated with other sequencing schemes (such as shared sequencer) other than the current single sequencer

