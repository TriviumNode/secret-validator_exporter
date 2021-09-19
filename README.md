# secret-validator_exporter
Prometheus exporter for Secret Network Validators


## Introduction
This exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter (localhost:26660), and other specific information monitoring purposes

## Usage

`go build`

`.\secret-validator_exporter \path\containing\configtoml\`

## Docker
```
docker run --name validator-exporter \
--publish 26661:26661 \
--env LCD_URL="http://172.17.0.1:1317" \
--env RPC_URL="172.17.0.1:26657" \
--env OPER_ADDR="cosmosvaloper1ahawe276d250zpxt0xgpfg63ymmu63a0svuvgw" \
--env PREFIX="cosmos" \
xiphiar/validator-exporter:1.0
```
