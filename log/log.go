package log

import "github.com/ccb1900/go-micro/log/zap"

type Log interface {
	GetName() string
	IsTraceEnabled() bool
	IsWarnEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsErrorEnabled() bool
	Trace(err ...error)
	Debug(err ...error)
	Warn(err ...error)
	Info(err ...error)
	Error(err ...error)
}

func Get(driver string) Log {
	return zap.New()
}
func Default() Log {
	return Get("zap")
}
