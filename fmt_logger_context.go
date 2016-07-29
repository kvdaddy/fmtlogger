package fmtlogger

type FmtLoggerContext struct {
	logger  FmtLogger
	keyvals []interface{}
}

func (l *FmtLoggerContext) Log(keyvals ...interface{}) error {
	kvs := append(l.keyvals, keyvals...)
	return l.logger.Log(kvs...)
}

func (l *FmtLoggerContext) Follow(keyvals ...interface{}) *FmtLoggerContext {
	kvs := append(l.keyvals, keyvals...)

	return &FmtLoggerContext{
		logger:  l.logger,
		keyvals: kvs,
	}
}

func NewContext(logger FmtLogger) *FmtLoggerContext {
	if c, ok := logger.(*FmtLoggerContext); ok {
		return c
	}
	return &FmtLoggerContext{logger: logger}
}
