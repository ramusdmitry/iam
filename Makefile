GOBASE = ${shell pwd}
GOBIN = ${GOBASE}/bin
PROJECT_NAME = iam
PROTO_API_DIR = api
PROTO_OUT_DIR = pkg/auth-api

.PHONY: gen/go
gen/go:
	go generate ./...

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o ${GOBIN}/${PROJECT_NAME} ./cmd/${PROJECT_NAME}/main.go || exit 1

.PHONY: gen/proto
gen/proto:
	rm -rf ${PROTO_OUT_DIR}
	mkdir -p ${PROTO_OUT_DIR}
	protoc -I ${PROTO_API_DIR} ./${PROTO_API_DIR}/*.proto \
      --go_out=${PROTO_OUT_DIR} --go_opt=paths=source_relative \
      --go-grpc_out=${PROTO_OUT_DIR} --go-grpc_opt=paths=source_relative
