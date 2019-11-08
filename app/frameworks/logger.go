package frameworks

import (
	"fmt"
	"io"
	"log"

	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
)

type Logger struct {
	I *log.Logger
	D *log.Logger
	E *log.Logger
}

func NewLogger(o, e io.Writer) usecases.Logger {
	return &Logger{
		I: log.New(o, "[INFO] ", log.LstdFlags),
		D: log.New(o, "[DEBUG] ", log.Llongfile|log.Ldate|log.Lmicroseconds),
		E: log.New(e, "[ERROR] ", log.Llongfile|log.Ldate|log.Lmicroseconds),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.I.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.I.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.D.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.D.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.E.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.E.Output(2, fmt.Sprintf(format, v...))
}
