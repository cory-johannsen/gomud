GITROOT := $(shell git rev-parse --show-toplevel)
GRPC_CONF=$(GITROOT)/server/protos/mud.yaml

.PHONY: clean generate build

build: generate wire
	mkdir -p $(GITROOT)/bin
	cd $(GITROOT)/server && \
	go build -o $(GITROOT)/bin/gunchete ./cmd

clean:
	rm -rf $(GITROOT)/bin
	rm -rf $(GITROOT)/server/dependencies
	rm -rf $(GITROOT)/server/generated

clone:
	mkdir -p $(GITROOT)/server/dependencies
	if [ -d $(GITROOT)/server/dependencies/googleapis ]; then \
		echo "googleapis already cloned"; \
	else \
		echo "Cloning googleapis..."; \
		cd $(GITROOT)/server/dependencies && \
		git clone https://github.com/googleapis/googleapis.git; \
	fi

generate: clone
	mkdir -p $(GITROOT)/server/generated
	cd server && \
	protoc -I protos \
		-I$(GITROOT)/server/dependencies/googleapis \
		--go_out=$(GITROOT)/server/generated \
		--go-grpc_out=$(GITROOT)/server/generated \
		--grpc-gateway_out=logtostderr=true,grpc_api_configuration=${GRPC_CONF}:$(GITROOT)/server/generated \
		--swagger_out=logtostderr=true,grpc_api_configuration=${GRPC_CONF}:$(GITROOT)/server/generated \
		protos/*.proto

wire:
	cd $(GITROOT)/server/internal/domain/effect && wire
	cd $(GITROOT)/server/cmd && wire
