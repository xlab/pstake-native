#!/bin/bash

PERSISTENCE_PORT="transfer"
PERSISTENCE_CHANNEL="channel-1"

set -o errexit -o nounset -o pipefail -eu

echo "## Check balances"
persistenceCore q bank balances $(persistenceCore keys show val1 -a)
persistenceCore q bank balances $(persistenceCore keys show val2 -a)
persistenceCore q bank balances $(persistenceCore keys show val3 -a)
persistenceCore q bank balances $(persistenceCore keys show val4 -a)

echo "## IBC Transfer atom from gaia to persistence"
echo "### Transfer 1 atom from gaia:val1 to persistence:val1"
gaiad tx ibc-transfer transfer "$PERSISTENCE_PORT" "$PERSISTENCE_CHANNEL" \
  $(persistenceCore keys show val1 -a) 1000000uatom \
  --from val1 --gas auto --gas-adjustment 1.2 -y --keyring-backend test > /dev/null
echo "### Transfer 10 atom from gaia:val2 to persistence:val4"
gaiad tx ibc-transfer transfer "$PERSISTENCE_PORT" "$PERSISTENCE_CHANNEL" \
  $(persistenceCore keys show val4 -a) 1000000uatom \
  --from val2 --gas auto --gas-adjustment 1.2 -y --keyring-backend test > /dev/null

echo "## Sleep for a bit to let ibc-transfer happen"
sleep 4

echo "## Check balances post upgrade transfer"
persistenceCore q bank balances $(persistenceCore keys show val1 -a)
persistenceCore q bank balances $(persistenceCore keys show val2 -a)
persistenceCore q bank balances $(persistenceCore keys show val3 -a)
persistenceCore q bank balances $(persistenceCore keys show val4 -a)
