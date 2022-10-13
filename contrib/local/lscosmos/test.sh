#!/bin/bash

CHAIN_BIN="${CHAIN_BIN:=persistenceCore}"
PERSISTENCE_PORT="transfer"
PERSISTENCE_CHANNEL="channel-1"

set -eu

IBC_DENOM=$($CHAIN_BIN q bank total | jq -r ".supply[] | select(.denom | startswith(\"ibc/\")) | .denom")

echo "## Check balances post upgrade transfer"
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val1 -a)
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val2 -a)
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val3 -a)
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val4 -a)

echo "## Perform liquid staking txn"
$CHAIN_BIN tx lscosmos liquid-stake 10000000$IBC_DENOM --from val3 -y
