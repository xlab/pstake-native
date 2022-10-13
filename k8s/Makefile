PROCESS=persistenceCore

all: delete apply

apply:
	kubectl apply -k $(PROCESS)

delete:
	kubectl delete -k $(PROCESS)

apply-all:
	$(MAKE) apply PROCESS=persistenceCore
	$(MAKE) apply PROCESS=gaia
	$(MAKE) apply PROCESS=relayer
	$(MAKE) apply PROCESS=explorer
	#$(MAKE) apply PROCESS=persistenceCore

delete-all:
	-$(MAKE) delete PROCESS=persistenceCore
	-$(MAKE) delete PROCESS=gaia
	-$(MAKE) delete PROCESS=relayer
	-$(MAKE) apply PROCESS=explorer

# debug:
    # port-forward
    # init -genesis
    # get seeds - by calling exposer
    # get genesis and paste in .$BINARY folder
    #
