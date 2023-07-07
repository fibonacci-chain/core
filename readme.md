# Fibonacci Chain

![FIBOLOGO.png](docs%2FFIBOLOGO.png)

Fibonacci is a forward-looking, high-performance public blockchain with the ability to be composable and iterative, and it is fully compatible with both EVM and WASM.

# fbchain
**fbchain** is a blockchain built using Cosmos SDK and Tendermint. It is built using the Cosmos SDK and Tendermint core

Leveraging its highly scalable underlying framework, Fibonacci is dedicated to building a customized SocialFi ecosystem for the social sector and creator economy.

# Documentation
For the most up to date documentation please visit https://fibochain.org

# Fibonacci Chain Ecosystem
Fibonacci Chain Network is an L1 blockchain with a built-in on-chain orderbook that allows smart contracts easy access to shared liquidity. Fibonacci Chain architecture enables composable apps that maintain modularity.

Fibonacci Chain Network serves as the matching core of the ecosystem, offering superior reliability and ultra-high transaction speed to ecosystem partners, each with their own functionality and user experience. Anyone can create a DeFi application that leverages Fibonacci Chain's liquidity and the entire ecosystem benefits.

Developers, traders, and users can all connect to Fibonacci Chain as ecosystem partners benefiting from shared liquidity and decentralized financial primitives.

# Testnet
## Get started
**How to validate on the Fibonacci Chain Testnet**
*This is the Fibonacci Chain Testnet-1 (fibonacci-testnet-1)*

## Hardware Requirements
**Minimum**
* 32 GB RAM
* 1 TB NVME SSD
* 16 Cores (modern CPU's)

## Operating System 

> Linux (x86_64) or Linux (amd64) Recommended Arch Linux

**Dependencies**
> Prerequisite: go1.18+ required.
* Arch Linux: `pacman -S go`
* Ubuntu: `sudo snap install go --classic`

> Prerequisite: git. 
* Arch Linux: `pacman -S git`
* Ubuntu: `sudo apt-get install git`

> Optional requirement: GNU make. 
* Arch Linux: `pacman -S make`
* Ubuntu: `sudo apt-get install make`

## Fibonacci Chaind Installation Steps

**Clone git repository**

```bash
git clone https://github.com/fibonacci-chain/core
cd core
git checkout origin/1.0.1beta-upgrade
make install
mv $HOME/go/bin/fbchaind /usr/bin/
```
**Generate keys**

* `fbchaind keys add [key_name]`

* `fbchaind keys add [key_name] --recover` to regenerate keys with your mnemonic

* `fbchaind keys add [key_name] --ledger` to generate keys with ledger device

## Validator setup instructions

* Install fbchaind binary

* Initialize node: `fbchaind init <moniker> --chain-id fbc-testnet-1`

* Download the Genesis file: `wget http://xxx -P $HOME/.fbchaind/config/`
 
* Edit the minimum-gas-prices in ${HOME}/.fbchaind/config/app.toml: `sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.01ufibo"/g' $HOME/.fbchaind/config/app.toml`

* Start fbchaind by creating a systemd service to run the node in the background
`nano /etc/systemd/system/fbchaind.service`
> Copy and paste the following text into your service file. Be sure to edit as you see fit.

```bash
[Unit]
Description=Fibonacci Chain-Network Node
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/
ExecStart=/root/go/bin/fbchaind start
Restart=on-failure
StartLimitInterval=0
RestartSec=3
LimitNOFILE=65535
LimitMEMLOCK=209715200

[Install]
WantedBy=multi-user.target
```
## Start the node

**Start fbchaind on Linux**

* Reload the service files: `sudo systemctl daemon-reload` 
* Create the symlinlk: `sudo systemctl enable fbchaind.service` 
* Start the node sudo: `systemctl start fbchaind && journalctl -u fbchaind -f`

**Start a chain on 4 node docker cluster**

* Start local 4 node cluster: `make docker-cluster-start`
* SSH into a docker container: `docker exec -it [container_name] /bin/bash`
* Stop local 4 node cluster: `make docker-cluster-stop`

### Create Validator Transaction
```bash
fbchaind tx staking create-validator \
--from {{KEY_NAME}} \
--chain-id  \
--moniker="<VALIDATOR_NAME>" \
--commission-max-change-rate=0.01 \
--commission-max-rate=1.0 \
--commission-rate=0.05 \
--details="<description>" \
--security-contact="<contact_information>" \
--website="<your_website>" \
--pubkey $(fbchaind tendermint show-validator) \
--min-self-delegation="1" \
--amount <token delegation>ufibo \
--node localhost:26657
```
# Build with Us!
If you are interested in building with Fibonacci Chain Network: 
Email us at fibonacci77777@gmail.com 
DM us on Twitter https://twitter.com/FIBOGlobal ChainNetwork
