simdcomet tx upgrade software-upgrade rollup-migrate --from mykey --upgrade-height 15 --upgrade-info "migration" --chain-id simdcomet-1 --keyring-backend test --yes --title upgrade-test --summary "migration" --no-checksum-required --no-validate --chain-id simdcomet-testnet-1 --deposit 10000000000stake

sleep 10

simdcomet tx gov vote 1 yes --from mykey --chain-id simdcomet-testnet-1  --keyring-backend test -y 