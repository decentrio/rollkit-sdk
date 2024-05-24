---
parent: Instructions
---

# Migration from CometBFT chain to Rollup

## Overview

This document outlines the steps required to migrate a CometBFT chain to Rollup.

## Prerequisites

Before starting the migration process, make sure your chain have the following:

- SDK verion >= v0.50.0
- CometBFT version >= v0.38.0

## Migration Steps

### 1. Prepare upgrade

- Integrate rollkit-sdk to your chain following [this guideline](./integration.md)

- Integrate `RappDK` upgrade handler to your chain. Example upgrade handler could be found here [upgrade handler](https://github.com/decentrio/rollkit-sdk/blob/01103c74832314f77b3a9271d18c33d393bc0529/simapp/upgrade/upgrade.go#L28).

**Notice: make sure you use correct sequencer pubkey address in your upgrade handler. You need to set the new sequencer pubkey manually in [your upgrade handler logic](https://github.com/decentrio/rollkit-sdk/blob/01103c74832314f77b3a9271d18c33d393bc0529/simapp/upgrade/upgrade.go#L30)**

### 2. Submit Upgrade proposal

- The chain create an upgrade proposal to include the `x/sequencer` module and upgrade handler above.

### 3. Migrate data

When the `upgrade_height` is reached, before run new binary, all node provider and validator need to migrate their node data by using this command:

```
[app_name] rollup-migration --home [node_dir]
```

And start the chain by using new binary. Now your chain is a rollup chain !

```bash
5:42PM INF Using pending block height=15 module=BlockManager
5:42PM INF starting gRPC server... address=127.0.0.1:9290 module=grpc-server
5:42PM INF applying upgrade "rollup-migrate" at height: 15 module=x/upgrade
5:42PM INF adding a new module: sequencer module=server
5:42PM INF Rollup changeover complete - you are now a rollup chain! module=x/sequencer
5:42PM INF finalized block block_app_hash=7801089B60A68C939C2F9017E70EA8E2F53B4772394322D1FEADFEBFC64491F1 height=15 module=BlockManager num_txs_res=0 num_val_updates=2
5:42PM INF executed block app_hash=7801089B60A68C939C2F9017E70EA8E2F53B4772394322D1FEADFEBFC64491F1 height=15 module=BlockManager
```

## Conclusion

By following the above steps, you can successfully migrate a chain from CometBFT to Rollup. Make sure to carefully review and test each step to ensure a smooth migration process.
