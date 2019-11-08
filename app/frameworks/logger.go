package frameworks

import (
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
	l.I.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.I.Printf(format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.D.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.E.Printf(format, v...)
}
