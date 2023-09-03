package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/meeron/honey-badger/config"
)

type Logger struct {
	logger *log.Logger
	src    string
}

var fileLogger *log.Logger
var loggers = make(map[string]*Logger)

func Init() error {
	config := config.Get().Logger

	os.Mkdir(config.Dir, 0777)
	logFile := fmt.Sprintf("%s.log", time.Now().Format("20060102"))

	f, err := os.OpenFile(path.Join(config.Dir, logFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	fileLogger = log.New(f, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return nil
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logMsg("ERROR", format, v...)
}

func (l *Logger) Error(err error) {
	l.logMsg("ERROR", err.Error())
}

func (l *Logger) Warningf(string, ...interface{}) {}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logMsg("INFO", format, v...)
}

func (l *Logger) Debugf(string, ...interface{}) {}

func (l *Logger) logMsg(level string, format string, v ...any) {
	l.logger.Printf("%s %s: %s", l.src, level, fmt.Sprintf(format, v...))
}

func Get(source string) *Logger {
	if l, ok := loggers[source]; ok {
		return l
	}

	loggers[source] = &Logger{
		logger: fileLogger,
		src:    source,
	}

	return loggers[source]
}

func Default() *Logger {
	return Get("default")
}

func Server() *Logger {
	return Get("server")
}

func Badger() *Logger {
	return Get("badger")
}
