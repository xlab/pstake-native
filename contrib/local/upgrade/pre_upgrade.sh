#!/bin/bash

PERSISTENCE_PORT="transfer"
PERSISTENCE_CHANNEL="channel-1"

set -o errexit -o nounset -o pipefail -eu

echo "## IBC Transfer atom from gaia to persistence"
echo "### Transfer 1 atom from gaia:val1 to persistence:val1"
gaiad tx ibc-transfer transfer "$PERSISTENCE_PORT" "$PERSISTENCE_CHANNEL" \
  $(persistenceCore keys show val1 -a) 1000000uatom \
  --from val1 --gas auto --gas-adjustment 1.2 -y --keyring-backend test > /dev/null
echo "### Transfer 15 atom from gaia:val1 to persistence:val2"
gaiad tx ibc-transfer transfer "$PERSISTENCE_PORT" "$PERSISTENCE_CHANNEL" \
  $(persistenceCore keys show val2 -a) 15000000uatom \
  --from val1 --gas auto --gas-adjustment 1.2 -y > /dev/null
echo "### Transfer 150 atom from gaia:val2 to persistence:val2"
gaiad tx ibc-transfer transfer "$PERSISTENCE_PORT" "$PERSISTENCE_CHANNEL" \
  $(persistenceCore keys show val3 -a) 150000000uatom \
  --from val2 --gas auto --gas-adjustment 1.2 -y > /dev/null

echo "## Sleep for a bit to let ibc-transfer happen"
sleep 4

echo "## Check balances"
persistenceCore q bank balances $(persistenceCore keys show val1 -a)
persistenceCore q bank balances $(persistenceCore keys show val2 -a)
persistenceCore q bank balances $(persistenceCore keys show val3 -a)

## Test IBC states
# Transfer atom from gaia to persistencecore couple of addresses
# Transfer xprt to gaia couple of addresses
# Check ibc channel/connection/client state
# > Upgrade to v4
# Check balance of ibc atom on persistencecore
# Transfer new tokens from gaia to persistence
# Check ibc channel/connection/client state

## Test LScosmos module testing
# > Upgrade to v4
# Jumpstart tx to enable ls cosmos module, read json from scripts: jump-start
# Perform liquid staking
