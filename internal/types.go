package internal

type FileHeader struct {
	// Package is the go package name
	Package string
	// LoggerEnabled is true if logger is enabled
	LoggerEnabled bool
}

type Handlers struct {
	// Hanlders are handlers(methods) under a single service
	Handlers []*Handler
	// Name is the struct name of the collection of handlers
	Name string
	// GrpcServerName is the server name that protoc-gen-go-grpc will generate
	// for the protocol buffer service
	GrpcServerName string
	// LoggerEnabled is true if logger is enabled
	LoggerEnabled bool
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
	// LoggerEnabled is true if logger is enabled
	LoggerEnabled bool
}
