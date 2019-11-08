package frameworks

import (
	"io"
	"log"

	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
)

type Logger struct {
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
}

func NewLogger(o, e io.Writer) usecases.Logger {
	return &Logger{
		Info:  log.New(o, "[INFO] ", log.Lshortfile|log.LstdFlags),
		Debug: log.New(o, "[DEBUG] ", log.Llongfile|log.Ldate|log.Lmicroseconds),
		Error: log.New(e, "[ERROR] ", log.Llongfile|log.Ldate|log.Lmicroseconds),
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error.Printf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info.Printf(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug.Printf(format, args...)
}
