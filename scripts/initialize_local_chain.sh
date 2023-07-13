#!/bin/bash
# require success for commands
set -e

# set key name
keyname=admin
CHAIN_ID=fbchain_9001-9001
#docker stop jaeger
#docker rm jaeger
#docker run -d --name jaeger \
#  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
#  -p 5775:5775/udp \
#  -p 6831:6831/udp \
#  -p 6832:6832/udp \
#  -p 5778:5778 \
#  -p 16686:16686 \
#  -p 14250:14250 \
#  -p 14268:14268 \
#  -p 14269:14269 \
#  -p 9411:9411 \
#  jaegertracing/all-in-one:1.33
# clean up old chain directory
rm -rf ~/.fbchaind
echo "Building..."
#install fbchaind
make install

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    cp ./build/fbchaind  ~/go/bin/fbchaind
    echo "fbchaind already"
fi

# initialize chain with chain ID and add the first key
~/go/bin/fbchaind init demo --chain-id $CHAIN_ID
~/go/bin/fbchaind keys add $keyname --keyring-backend test
# add the key as a genesis account with massive balances of several different tokens
~/go/bin/fbchaind add-genesis-account $(~/go/bin/fbchaind keys show $keyname -a --keyring-backend test) 1000000000000000000000000000ufibo,100000000000000000000uusdc,100000000000000000000uatom
# gentx for account
~/go/bin/fbchaind gentx $keyname 70000000000000000000ufibo --chain-id $CHAIN_ID --keyring-backend test

# update config to run as a validator, add validator information to genesis file
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  sed -i  's/mode = "full"/mode = "validator"/g' $HOME/.fbchaind/config/config.toml
  sed -i  's/indexer = \["null"\]/indexer = \["kv"\]/g' $HOME/.fbchaind/config/config.toml
elif [[ "$OSTYPE" == "darwin"* ]]; then
  sed -i '' 's/mode = "full"/mode = "validator"/g' $HOME/.fbchaind/config/config.toml
  sed -i '' 's/indexer = \["null"\]/indexer = \["kv"\]/g' $HOME/.fbchaind/config/config.toml
else
  echo "OS not supported"
fi

KEY=$(jq '.pub_key' ~/.fbchaind/config/priv_validator_key.json -c)
jq '.validators = [{}]' ~/.fbchaind/config/genesis.json > ~/.fbchaind/config/tmp_genesis.json
jq '.validators[0] += {"power":"70000000000000"}' ~/.fbchaind/config/tmp_genesis.json > ~/.fbchaind/config/tmp_genesis_2.json
jq '.validators[0] += {"pub_key":'$KEY'}' ~/.fbchaind/config/tmp_genesis_2.json > ~/.fbchaind/config/tmp_genesis_3.json
mv ~/.fbchaind/config/tmp_genesis_3.json ~/.fbchaind/config/genesis.json && rm ~/.fbchaind/config/tmp_genesis.json && rm ~/.fbchaind/config/tmp_genesis_2.json

echo "Creating Accounts"
# create 10 test accounts + fund them
python3  loadtest/scripts/populate_genesis_accounts.py 10 loc

 ~/go/bin/fbchaind collect-gentxs
if [[ "$OSTYPE" == "linux-gnu" ]]; then
    # update some params in genesis file for easier use of the chain localls (make gov props faster)
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["max_deposit_period"]="300s"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="30s"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq --arg start_date "$(date +"%Y-%m-%d")" --arg end_date "$(date -d "+3days" +"%Y-%m-%d")" '.app_state["mint"]["params"]["token_release_schedule"]=[{"start_date": $start_date, "end_date": $end_date, "token_release_amount": "999999999999"}]' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq --arg start_date "$(date -d "+3days" +"%Y-%m-%d")" --arg end_date "$(date -d "+5days" +"%Y-%m-%d")" '.app_state["mint"]["params"]["token_release_schedule"] += [{"start_date": $start_date, "end_date": $end_date, "token_release_amount": "999999999999"}]' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["gov"]["voting_params"]["expedited_voting_period"]="10s"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["oracle"]["params"]["vote_period"]="1"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["oracle"]["params"]["whitelist"]=[{"name": "ueth"},{"name": "ubtc"},{"name": "uusdc"},{"name": "uusdt"}]' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["distribution"]["params"]["community_tax"]="0.000000000000000000"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json

elif [[ "$OSTYPE" == "darwin" ]]; then
    # update some params in genesis file for easier use of the chain localls (make gov props faster)
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["max_deposit_period"]="300s"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="30s"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq --arg start_date "$(date +"%Y-%m-%d")" --arg end_date "$(date -v+3d +"%Y-%m-%d")" '.app_state["mint"]["params"]["token_release_schedule"]=[{"start_date": $start_date, "end_date": $end_date, "token_release_amount": "999999999999"}]' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq --arg start_date "$(date -v+3d +"%Y-%m-%d")" --arg end_date "$(date -v+5d +"%Y-%m-%d")" '.app_state["mint"]["params"]["token_release_schedule"] += [{"start_date": $start_date, "end_date": $end_date, "token_release_amount": "999999999999"}]' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["gov"]["voting_params"]["expedited_voting_period"]="10s"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["oracle"]["params"]["vote_period"]="1"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["oracle"]["params"]["whitelist"]=[{"name": "ueth"},{"name": "ubtc"},{"name": "uusdc"},{"name": "uusdt"}]' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
    cat ~/.fbchaind/config/genesis.json | jq '.app_state["distribution"]["params"]["community_tax"]="0.000000000000000000"' > ~/.fbchaind/config/tmp_genesis.json && mv ~/.fbchaind/config/tmp_genesis.json ~/.fbchaind/config/genesis.json
# elif [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys" || "$OSTYPE" == "win32" || "$OSTYPE" == "win64" ]]; then
#     echo "This is Windows"
else
    echo "Unknown operating system"
fi

# set block time to 2s
if [ ! -z "$1" ]; then
  CONFIG_PATH="$1"
else
  CONFIG_PATH="$HOME/.fbchaind/config/config.toml"
fi

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  sed -i 's/timeout_prevote =.*/timeout_prevote = "2000ms"/g' $CONFIG_PATH
  sed -i 's/timeout_precommit =.*/timeout_precommit = "2000ms"/g' $CONFIG_PATH
  sed -i 's/timeout_commit =.*/timeout_commit = "2000ms"/g' $CONFIG_PATH
  sed -i 's/skip_timeout_commit =.*/skip_timeout_commit = false/g' $CONFIG_PATH
elif [[ "$OSTYPE" == "darwin"* ]]; then
  sed -i '' 's/unsafe-propose-timeout-override =.*/unsafe-propose-timeout-override = "2s"/g' $CONFIG_PATH
  sed -i '' 's/unsafe-propose-timeout-delta-override =.*/unsafe-propose-timeout-delta-override = "2s"/g' $CONFIG_PATH
  sed -i '' 's/unsafe-vote-timeout-override =.*/unsafe-vote-timeout-override = "2s"/g' $CONFIG_PATH
  sed -i '' 's/unsafe-vote-timeout-delta-override =.*/unsafe-vote-timeout-delta-override = "2s"/g' $CONFIG_PATH
  sed -i '' 's/unsafe-commit-timeout-override =.*/unsafe-commit-timeout-override = "2s"/g' $CONFIG_PATH
else
  printf "Platform not supported, please ensure that the following values are set in your config.toml:\n"
  printf "###         Consensus Configuration Options         ###\n"
  printf "\t timeout_prevote = \"2000ms\"\n"
  printf "\t timeout_precommit = \"2000ms\"\n"
  printf "\t timeout_commit = \"2000ms\"\n"
  printf "\t skip_timeout_commit = false\n"
  exit 1
fi

~/go/bin/fbchaind config keyring-backend test

# start the chain with log tracing
# --evm.tracer json
GORACE="log_path=/tmp/race/fbchaind_race" ~/go/bin/fbchaind start --trace --chain-id ${CHAIN_ID}
