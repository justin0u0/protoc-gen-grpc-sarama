PATH := $(CURDIR)/bin:$(PATH)

.PHONY: dc.build
dc.build:
	docker compose run --rm build

.PHONY: build
build:
	mkdir -p ./bin/
	go build -o ./bin/protoc-gen-grpc-sarama ./main.go

.PHONY: dc.generate
dc.generate:
	docker compose run --rm generate

.PHONY: generate
generate: bin/protoc-gen-go
	protoc \
	-I . \
	--go_out=paths=source_relative:. \
	./proto/*.proto

.PHONY: dc.example
dc.example:
	docker compose run --rm example

.PHONY: example
example: bin/protoc-gen-go bin/protoc-gen-go-grpc
	protoc \
	-I . \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	--grpc-sarama_out=paths=source_relative:. \
	./example/*.proto

bin/protoc-gen-go: go.mod
	go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

bin/protoc-gen-go-grpc: go.mod
	go build -o $@ google.golang.org/grpc/cmd/protoc-gen-go-grpc
