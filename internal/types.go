package internal

type Handlers struct {
	// Hanlders are handlers(methods) under a single service
	Handlers []*Handler
	// GrpcServerName is the server name that protoc-gen-go-grpc will generate
	// for the protocol buffer service
	GrpcServerName string
}

type Handler struct {
	// Name is the struct name of each handler
	Name string
	// GrpcServerName is the server name that protoc-gen-go-grpc will generate
	// for the protocol buffer service
	GrpcServerName string
	// Method is the method name
	Method string
	// Input is the method input name
	Input string
}
