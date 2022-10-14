#!/bin/bash

CHAIN_BIN="${CHAIN_BIN:=persistenceCore}"
PERSISTENCE_PORT="transfer"
PERSISTENCE_CHANNEL="channel-1"

set -eu

echo "## Check balances post upgrade transfer"
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val1 -a)
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val2 -a)
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val3 -a)
$CHAIN_BIN q bank balances $($CHAIN_BIN keys show val4 -a)

echo "## Submit proposal to set persistence:val4 as jumpstart address"

cat << EOF > /tmp/update_proposal.json
{
  "title": "register host chain proposal",
  "description": "this proposal register host chain params in the chain",
  "pstake_fee_address": "$($CHAIN_BIN keys show val4 -a)",
  "deposit": "1000000uxprt"
}
EOF

echo "### Checking if proposal already passed"
LOOKUP_PROPOSAL=$($CHAIN_BIN q gov proposals | jq "last(.proposals[] | select(.content.title == $(cat /tmp/update_proposal.json | jq .title)))")

if [[ -z $LOOKUP_PROPOSAL ]] || [[ $LOOKUP_PROPOSAL == "null" ]]; then
  echo "### Submit proposal from val4"
  RESP=$($CHAIN_BIN tx gov submit-proposal pstake-lscosmos-change-pstake-fee-address \
    /tmp/update_proposal.json -y --from val4 --fees 2000uxprt --gas auto --gas-adjustment 1.5 -b block -o json)
  #echo $RESP | jq
  PROPOSAL_ID=$(echo "$RESP" | jq -r '.logs[0].events[] | select(.type == "submit_proposal") | .attributes[] | select(.key == "proposal_id") | .value')
  echo "* PROPOSAL_ID: $PROPOSAL_ID"

  echo "### Query proposal prevote"
  $CHAIN_BIN q gov proposal $PROPOSAL_ID -o json > /dev/null

  echo "### Vote proposal"
  $CHAIN_BIN tx gov vote $PROPOSAL_ID yes --from val4 --yes \
      --fees 200uxprt --gas auto --gas-adjustment 1.5 -b block -o json > /dev/null
  $CHAIN_BIN tx gov vote $PROPOSAL_ID yes --from val2 --yes \
      --fees 200uxprt --gas auto --gas-adjustment 1.5 -b block -o json > /dev/null

  echo "###Proposal voting period"
  sleep 40
  echo "### Query proposal postvote"
  $CHAIN_BIN q gov proposal $PROPOSAL_ID -o json | jq ".content"
else
  echo "* Looks like the proposal already submitted... skipping proposal submission"
  echo "PROPOSAL_ID: $(echo $LOOKUP_PROPOSAL | jq ".proposal_id")"
fi

echo "### Check set fee address"
$CHAIN_BIN q lscosmos host-chain-params | jq ".host_chain_params.pstake_params.pstake_fee_address"

echo "## Creating jumpstart txn on ls-cosmos module"
CHANNEL_INFO=$($CHAIN_BIN q ibc channel channels  | jq '.channels[] | select(.state == "STATE_OPEN") | select(.port_id == "transfer")')
echo "### Channel info"
echo $CHANNEL_INFO | jq

if [[ -z $CHANNEL_INFO ]]; then
  echo "No open transfer port and connection.... exiting";
  exit 1;
fi

cat << EOF > /tmp/jumpstart.json
{
  "chain_id": "$(gaiad status 2>&1 | jq -r ".NodeInfo.network")",
  "connection_id": "$(echo $CHANNEL_INFO | jq -r '.connection_hops[0]')",
  "transfer_channel": "$(echo $CHANNEL_INFO | jq -r '.channel_id')",
  "transfer_port": "$(echo $CHANNEL_INFO | jq -r '.port_id')",
  "base_denom": "uatom",
  "mint_denom": "stk/uatom",
  "min_deposit": "1",
  "allow_listed_validators": {
    "allow_listed_validators": [
      {
        "validator_address": "$(gaiad keys show val1 -a --bech val)",
        "target_weight": "1"
      }
    ]
  },
  "pstake_params": {
    "pstake_deposit_fee": "0.00",
    "pstake_restake_fee": "0.05",
    "pstake_unstake_fee": "0.00",
    "pstake_redemption_fee": "0.1",
    "pstake_fee_address": "$($CHAIN_BIN keys show val4 -a)"
  },
  "host_accounts": {
    "delegator_account_owner_i_d": "lscosmos_pstake_delegation_account",
    "rewards_account_owner_i_d": "lscosmos_pstake_reward_account"
  }
}
EOF

echo "### Try to perform jumpstart from wrong address"
set +e
$CHAIN_BIN tx lscosmos jump-start /tmp/jumpstart.json \
  --from val1 -y --gas auto --gas-adjustment 1.5 2> /dev/null
if [ -$? -eq 0 ]; then
  echo "* Something went wrong... jumpstart worked with persistence:val1, set persistence:val4"
  echo "exiting..."
  exit 1
else
  echo "* lscosmo jumpstart txn did not with non fee address, as expected"
fi
set -e

set +e
echo "## Perform Jumpstart from correct address"
RESP=$($CHAIN_BIN tx lscosmos jump-start /tmp/jumpstart.json \
  --from val4 -y --gas auto --gas-adjustment 1.5 2>&1)

if [ -$? -eq 0 ]; then
  echo $RESP | jq ".raw_log";
elif [[ "$RESP" == *"Error:"* ]] && [[ "$RESP" == *"Module is already enabled"* ]]; then
  echo "* module ls-cosmos already enabled... continuing";
else
  echo $RESP;
  exit 1;
fi
set -e

echo "### Try to redo jumpstart"
set +e
$CHAIN_BIN tx lscosmos jump-start /tmp/jumpstart.json \
  --from val4 -y --gas auto --gas-adjustment 1.5 2> /dev/null
if [ -$? -eq 0 ]; then
  echo "* Something went wrong... jumpstart worked even it was enabled"
  echo "exiting..."
  exit 1
else
  echo "* lscosmo redo jumpstart txn did not work since enabled, as expected"
fi
set -e


