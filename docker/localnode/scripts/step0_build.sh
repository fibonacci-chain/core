#!/usr/bin/env sh

# Input parameters
NODE_ID=${ID:-0}
ARCH=$(uname -m)

# Build fbchaind
echo "Building fbchaind from local branch"
git config --global --add safe.directory /fibonacci-chain/core
LEDGER_ENABLED=false
make clean
make build-linux
make build-price-feeder-linux
mkdir -p build/generated
echo "DONE" > build/generated/build.complete
