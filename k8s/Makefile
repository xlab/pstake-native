PROCESS=persistenceCore

all: delete-all apply-all

apply:
	kubectl apply -k $(PROCESS)

delete:
	kubectl delete -k $(PROCESS)

apply-all:
	$(MAKE) apply PROCESS=persistenceCore
	$(MAKE) apply PROCESS=gaia
	$(MAKE) apply PROCESS=relayer
	$(MAKE) apply PROCESS=explorer

delete-all:
	-$(MAKE) delete PROCESS=persistenceCore
	-$(MAKE) delete PROCESS=gaia
	-$(MAKE) delete PROCESS=relayer
	-$(MAKE) delete PROCESS=explorer
