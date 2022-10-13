PERSISTENCE_KEYS="../../k8s/persistenceCore/configs/validators.json"
GAIA_KEYS="../../k8s/gaia/configs/validators.json"

setup-persistence-keys:
	CHAIN_BIN=persistenceCore CHAIN_ID=persistencecore-1 \
	DENOM=uxprt CHAIN_DATA_DIR=.persistenceCore NODE_PORT=26657 \
	VALIDATOR_CONFIG=$(PERSISTENCE_KEYS) \
	bash -e setup_keys.sh

clean-persistence-keys:
	rm -rf ~/.persistenceCore/keyring-test/

setup-gaia-keys:
	CHAIN_BIN=gaiad CHAIN_ID=gaiad-1 \
	DENOM=uatom CHAIN_DATA_DIR=.gaia NODE_PORT=27657 \
	VALIDATOR_CONFIG=$(GAIA_KEYS) \
	bash -e setup_keys.sh

clean-gaia-keys:
	rm -rf ~/.gaia/keyring-test/

setup-keys: setup-persistence-keys setup-gaia-keys

clean-keys: clean-persistence-keys clean-gaia-keys

## Upgrade commands

pre-upgade:
	bash -e upgrade/pre_upgrade.sh

.PHONY: upgrade
upgrade:
	bash -e upgrade/upgrade.sh

.PHONY: post-upgrade
post-upgrade:
	bash -e upgrade/post_upgrade.sh

## Lscosmos testing
setup-lscosmos:
	bash -e lscosmos/setup_lscosmos.sh

test-lscosmos:
	bash -e lscosmos/test.sh
