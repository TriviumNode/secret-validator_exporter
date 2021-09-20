## Introduction
This exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter (localhost:26660), and other specific information monitoring purposes

# secret-validator_exporter :satellite:
![CreatePlan](https://img.shields.io/badge/go-1.12.4%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)

## Docker Quick-Start
```
docker run --name validator-exporter \
--publish 26661:26661 \
--env LCD_URL="http://172.17.0.1:1317" \
--env RPC_URL="172.17.0.1:26657" \
--env OPER_ADDR="secretvaloper1ahawe276d250zpxt0xgpfg63ymmu63a0svuvgw" \
--env PREFIX="secret" \
xiphiar/validator-exporter:latest
```

## Metrics Available
> **Network**
- chainId: Name of the chain
- blockHeight: Height of the current block
- bondedTokens: Number of currently bonded SCRT
- notBondedTokens: Number of unbonded SCRT
- totalBondedTokens: Number of currently bonded & unbonded SCRT
- bondedRatio: Ratio of bonded tokens within the network

> **Validator Info**
- moniker: Name of the validator
- accountAddress: Validator's Account address
- consHexAddress: Validator's Consensus Hex address
- operatorAddress: Validator's Operator address
- validatorPubKey: Validator's Validator pubkey(```secretd tendermint show-validator```)
- votingPower: Decimal truncated Total voting power of the validator
- delegatorShares: Validator's total delegated tokens
- delegatorCount: Number of each unique delegators for a validator
- delegationRatio: Ratio of validator's bonded tokens to the network's total bonded tokens
- selfDelegationAmount: Self-bonded amount of the validator
- proposerPriorityValue: Number which represents the priority of the validator proposing in the next round
- proposerPriority: Rank of the proposerPriorityValue
- proposingStatus: Shows if the validator is the proposer or not in the current round(true: 1, false: 0)
- validatorCommitStatus: Confirms if the validator has committed in this round(true: 1, false: 0)
- commissionMaxChangeRate: Max range of commission rate whic hthe validator can change
- commissionMaxRate: The highest commission rate which the validator can charge
- commissionRate: Commission rate of the validator charged on delegators' rewards
- balances(uscrt): Wallet information of the validator which shows the balance
- commission(uscrt): Accumulated commission fee of the validator
- rewards(uscrt): Accumulated rewards of the validator
- minSelfDelegation(SCRT): The required minimum number of tokens which the validator must self-delegate
- jailed: Confirms if the validator is jailed or not(true: 1, false: 0)


## Grafana Example
 - Template: https://grafana.com/grafana/dashboards/10942/revisions
 
Can set alarms using the functions on Grafana (ex. Alarms if the validator fails to precommit or gets jailed)


## Build
```bash
go build
```

## Config
1. Modify to the appropriate RPC and REST server address
2. Modify the value of ```operatorAddr``` to the operator address of the validator you want to monitor.
3. You can change the service port(default: 26661)
```bash
# TOML Document for Cosmos-Validator Exporter(Pometheus & Grafana)

title = "TOML Document"

[Servers]
        [Servers.addr]
        rpc = "localhost:26657"
        rest = "localhost:1317"

[Validator]
operatorAddr = "secretvaloper14l0fp639yudfl46zauvv8rkzjgd4u0zk2aseys"

[Options]
listenPort = "26661"

```

## Start
  
```bash
./secret-validator_exporter {path containing config.toml}

// ex)
./secret-validator_exporter /data/secret/exporter/
```

## Use systemd service
  
```sh
# create user 'secret'
adduser secret

# Make log directory & file
sudo mkdir /var/log/userLog  
sudo touch /var/log/userLog/secret-validator_exporter.log  
sudo chown secret:secret /var/log/userLog/secret-validator_exporter.log

# Setup working directory
mkdir /home/secret/exporter
cp ./secret-validator_exporter /home/secret/exporter/
cp ./config.toml /home/secret/exporter/
sudo chown -R secret:secret /home/secret/exporter

# Create systemd service
sudo tee /etc/systemd/system/secret-validator_exporter.service > /dev/null <<EOF
[Unit]
Description=Secret Validator Exporter
After=network-online.target

[Service]
User=secret
WorkingDirectory=/home/secret/exporter
ExecStart=/home/secret/exporter/secret-validator_exporter \
        /home/secret/exporter
StandardOutput=file:/var/log/userLog/secret-validator_exporter.log
StandardError=file:/var/log/userLog/secret-validator_exporter.log
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl enable secret-validator_exporter.service
sudo systemctl restart secret-validator_exporter.service


## log
tail -f /var/log/userLog/secret-validator_exporter.log
```
