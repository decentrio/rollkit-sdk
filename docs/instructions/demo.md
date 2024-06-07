---
parent: Instructions
---

# Run the demo RappDK rollup


## Prerequisites

- Go version >= 1.21

## Install demo app

```bash
https://github.com/decentrio/rollkit-sdk
cd rollkit-sdk
make install
```

## Starting your rollup using rollkit-sdk

### Run the mock DA

```bash
wget https://github.com/decentrio/rollkit-sdk/blob/main/scripts/init-mock-da.sh
sh install-mock-da.sh
```

### Run the chain by using script

Download the required scripts

```bash
wget https://github.com/decentrio/rollkit-sdk/blob/main/scripts/init-local.sh
```

```bash
bash init-local.sh
```
