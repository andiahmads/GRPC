PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client
RM_F_CMD = rm -f
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

.DEFAULT_GOAL := help
.PHONY: hello-grpc product help
project := hello-grpc product

all: $(project) ## Generate Pbs and build

hello-grpc: $@ ## Generate Pbs and build for hello-grpc

$(project):
	@echo "Go package: ${PACKAGE}"
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto

clean_greet: ## Clean generated files for greet
	${RM_F_CMD} greet/${PROTO_DIR}/*.pb.go


hello-server:
	go run hello-grpc/server/main.go

hello-client:
	go run hello-grpc/client/main.go

grpc-test:
	evans --host localhost --port 2023 --reflection repl

product-server:
	go run product/cmd/main.go

product-client:
	go run product/client/main.go