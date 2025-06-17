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

clone-prepare:
	mkdir -p $(GITROOT)/server/dependencies

clone-google-apis:
	if [ -d $(GITROOT)/server/dependencies/googleapis ]; then \
		echo "googleapis already cloned"; \
	else \
		echo "Cloning googleapis..."; \
		cd $(GITROOT)/server/dependencies && \
		git clone https://github.com/googleapis/googleapis.git; \
	fi

clone-swagger-ui:
	if [ -d $(GITROOT)/server/dependencies/swagger-ui ]; then \
		echo "swagger-ui already cloned"; \
	else \
		echo "Cloning swagger-ui..."; \
		cd $(GITROOT)/server/dependencies && \
		wget https://github.com/swagger-api/swagger-ui/archive/refs/tags/v5.24.2.tar.gz && \
		tar -xzf v5.24.2.tar.gz && \
		mv swagger-ui-5.24.2 swagger-ui && \
		rm v5.24.2.tar.gz; \
	fi

clone: clone-prepare clone-google-apis clone-swagger-ui

generate-prepare:
	mkdir -p $(GITROOT)/server/generated

generate-proto:
	cd server && \
	protoc -I protos \
		-I$(GITROOT)/server/dependencies/googleapis \
		--go_out=$(GITROOT)/server/generated \
		--go-grpc_out=$(GITROOT)/server/generated \
		--grpc-gateway_out=logtostderr=true,grpc_api_configuration=${GRPC_CONF}:$(GITROOT)/server/generated \
		--swagger_out=logtostderr=true,grpc_api_configuration=${GRPC_CONF}:$(GITROOT)/server/generated \
		protos/*.proto

generate-swagger-ui:
	# Patch the generated swagger.json to use the correct base path
	mkdir -p $(GITROOT)/server/generated/swagger-ui && \
	cp $(GITROOT)/server/dependencies/swagger-ui/dist/* $(GITROOT)/server/generated/swagger-ui && \
	sed -i '' 's|^\([[:space:]]*url:[[:space:]]*\).*|\1\"/mud.swagger.json\",|' $(GITROOT)/server/generated/swagger-ui/swagger-initializer.js

generate: clone generate-prepare generate-proto generate-swagger-ui

wire:
	cd $(GITROOT)/server/internal/domain/effect && wire
	cd $(GITROOT)/server/cmd && wire
