# Sequencer module

Sequencer is the module that manage the actual valset (which is the sequencer).

It's the only module that can make `abci valset update` to rollkit.

Since the x/staking now only manage the governators, it should no longer meant to manage the actual valset anymore (separation of concerns). This mean we should have another module to handle the task so that we still keep abci semantic intact (enable abci valset update for rollkit).

The module will be easy to maintain as it only has to manage the one single sequencer.

Doing this we can cleanly solve the problem of dummy token requirement when initializing a rollkit because we now have the sequencer module init the actual valset (sequencer) instead of the staking module.

Further more, we can put useful features such as changing sequencer via gov or intergrate with shared sequencer.

