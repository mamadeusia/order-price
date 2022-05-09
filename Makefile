GOPATH:=$(shell go env GOPATH)


# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOPKG := $(.)

all : go-build-price go-build-order proto

# A valid GOPATH is required to use the `go get` command.
# If $GOPATH is not specified, $HOME/go will be used by default
# GOPATH := $(if $(GOPATH),$(GOPATH),~/go)
go-build-order:
	@echo "  >  Building order binary..."
	GOBIN=$(GOBIN) go build -o $(GOBIN)/order ./cmd/order



go-build-price:
	@echo "  >  Building price binary..."
	GOBIN=$(GOBIN) go build -o $(GOBIN)/price ./cmd/price


proto:
	mkdir -p server/grpc/pricegrpc

	protoc -I=server/grpc/pb \
	--go_out=server/grpc/pricegrpc \
	--go-grpc_out=server/grpc/pricegrpc \
	server/grpc/pb/price.proto