package main

import (
	"github.com/justin0u0/protoc-gen-grpc-sarama/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	g := &internal.Generator{}

	protogen.Options{}.Run(g.Generate)
}
