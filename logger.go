package musicflow

type Logger interface {
	Printf(format string, v ...interface{})
}

func WithLogger(log Logger) DialOption {
	return func(o *dialOptions) {
		o.logger = log
	}
}

type noopLogger struct{}

func (noopLogger) Printf(format string, v ...interface{}) {}
