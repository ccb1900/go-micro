package zap

type Log struct {
}

func (l *Log) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (l *Log) IsTraceEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *Log) IsWarnEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *Log) IsDebugEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *Log) IsInfoEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *Log) IsErrorEnabled() bool {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Trace(err ...error) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Debug(err ...error) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Warn(err ...error) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Info(err ...error) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Error(err ...error) {
	//TODO implement me
	panic("implement me")
}

func New() *Log {
	return &Log{}
}
