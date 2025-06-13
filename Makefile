GITROOT := $(shell git rev-parse --show-toplevel)

.PHONY: clean generate build

clean:
	rm -rf $(GITROOT)/bin
	rm -rf $(GITROOT)/server/generated

generate:
	mkdir -p $(GITROOT)/server/generated
	cd server && \
	protoc --go_out=$(GITROOT)/server/generated --go-grpc_out=$(GITROOT)/server/generated protos/*.proto

wire:
	cd $(GITROOT)/server/internal/domain/effect && wire
	cd $(GITROOT)/server/cmd && wire

build: generate wire
	mkdir -p $(GITROOT)/bin
	cd $(GITROOT)/server && \
	go build -o $(GITROOT)/bin/gunchete ./cmd