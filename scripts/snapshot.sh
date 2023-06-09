#!/bin/bash

mkdir $HOME/sei-snapshot
mkdir $HOME/key_backup
# Move priv_validator_state out so it isn't used by anyone else
mv $HOME/.fbchaind/data/priv_validator_state.json $HOME/key_backup
# Create backups
cd $HOME/sei-snapshot
tar -czf data.tar.gz -C $HOME/.fbchaind data/
tar -czf wasm.tar.gz -C $HOME/.fbchaind wasm/
echo "Data and Wasm snapshots created in $HOME/sei-snapshot"