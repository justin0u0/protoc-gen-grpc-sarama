# protoc-gen-grpc-sarama

The `protoc-gen-grpc-sarama` is a plugin of the Google protocol buffers compiler protoc. It reads protobuf service definitions and generates a [`ConsumerGroupHandler`](https://pkg.go.dev/github.com/Shopify/sarama#ConsumerGroupHandler) of the package [Shopify/sarama](https://github.com/Shopify/sarama).

See the [example protocol buffer file](example/example.proto) and the [generated file](example/example.pb.sarama.go) for more details.

For configuration, see [proto/sarama.proto](proto/sarama.proto) for more details.
