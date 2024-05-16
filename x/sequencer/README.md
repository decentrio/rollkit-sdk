# Sequencer module

Since the x/staking now only manage the governators, it should no longer meant to manage the actual valset anymore. This mean we should have another module to handle said task (separation of concerns) so that we still keep abci semantic intact (enable abci valset update for rollkit)

Sequencer is the module that manage the actual valset (the sequencer). It's the only module that can make `abci valset update` to rollkit. When initializing a rollkit, we now have the sequencer module init the actual valset (sequencer) instead of the staking module

This module can potentially be integrated with other sequencing schemes (such as shared sequencer) other than the current single sequencer
