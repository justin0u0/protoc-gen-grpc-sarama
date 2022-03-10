package saramakit

type Logger interface {
	// With carries a key-value pair data to the logger
	With(key, value string) Logger
	// Error prints an error log
	Error(msg string, err error)
}
