#!/usr/bin/env sh

NODE_ID=${ID:-0}
INVARIANT_CHECK_INTERVAL=${INVARIANT_CHECK_INTERVAL:-0}

LOG_DIR="build/generated/logs"
mkdir -p $LOG_DIR

# Starting sei chain
echo "Starting the fibo process for node $NODE_ID with invariant check interval=$INVARIANT_CHECK_INTERVAL..."
cp build/generated/genesis.json ~/.fbchaind/config/genesis.json
fibochain start --chain-id fbchaind_9001-9001 --inv-check-period ${INVARIANT_CHECK_INTERVAL} > "$LOG_DIR/fbchaind-$NODE_ID.log" 2>&1 &
echo "Node $NODE_ID fbchain is started now"
echo "Done" >> build/generated/launch.complete
