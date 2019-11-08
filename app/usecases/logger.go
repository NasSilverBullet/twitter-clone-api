package usecases

type Logger interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}
