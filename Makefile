BIN_FILE := "./bin/survey"
DOCKER_IMG="survey:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

gen:
	rm -f ./internal/server/grpc/api/api.pb.go
	rm -f ./internal/server/grpc/api/api_grpc.pb.go
	rm -f ./internal/server/grpc/api/api.pb.gw.go
	protoc --go-grpc_out=./internal/server/grpc/api api/*.proto
	protoc --go_out=./internal/server/grpc/api \
		--grpc-gateway_out=./internal/server/grpc/api \
		--openapiv2_out ./swagger/ \
		api/*.proto
	cp ./swagger/api/api.swagger.json ./swagger/swagger.json
	# --grpc-gateway_opt generate_unbound_methods=true - включение не аннотированных методов 

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-a -installsuffix cgo \
	-ldflags "$(LDFLAGS)" \
    -o ${BIN_FILE} ./cmd/survey/main.go ./cmd/survey/version.go

run: build
	$(BIN) --config=./cmd/survey/config.json

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f ./build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/... ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build run build-img run-img version test lint
