#!/bin/bash

CHAIN_BIN="${CHAIN_BIN:=persistenceCore}"
DENOM="${DENOM:=uxprt}"
CHAIN_DATA_DIR="${CHAIN_DATA_DIR:=.persistenceCore}"
CHAIN_ID="${CHAIN_ID:=persistencecore-1}"

set -o errexit -o nounset -o pipefail -eu

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
