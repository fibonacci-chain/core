#!/usr/bin/env sh

# Input parameters
ARCH=$(uname -m)

# Build fbchain
echo "Building fbchain from local branch"
git config --global --add safe.directory /fibonacci-chain/core
LEDGER_ENABLED=false
make install
mkdir -p build/generated
echo "DONE" > build/generated/build.complete
