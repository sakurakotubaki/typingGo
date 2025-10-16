.PHONY: run
GOMAIN ?= main.go

run:
	@echo "Running $(GOMAIN)..."
	go run $(GOMAIN)