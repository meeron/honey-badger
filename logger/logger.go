package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/meeron/honey-badger/config"
)

var current *log.Logger

func Init() error {
	if current != nil {
		return errors.New("logger already created")
	}

	config := config.Get().Logger

	os.Mkdir(config.Dir, 0777)
	logFile := fmt.Sprintf("%s.log", time.Now().Format("20060102"))

	f, err := os.OpenFile(path.Join(config.Dir, logFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	current = log.New(f, "HoneyBadger ", log.Ldate|log.Ltime|log.Lmicroseconds)
	return nil
}

func Info(format string, v ...any) {
	logMsg("INFO", format, v...)
}

func Warn(format string, v ...any) {
	logMsg("WARN", format, v...)
}

func Error(err error) {
	logMsg("ERROR", err.Error())
}

func logMsg(level string, format string, v ...any) {
	current.Printf("%s: %s", level, fmt.Sprintf(format, v...))
}
