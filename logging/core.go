package logging

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

type logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

var (
	l *logger
)

func InitialiseLogger() {
	l = &logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lmsgprefix),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lmsgprefix),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lmsgprefix),
	}
}

func getCallerInfo(frame int) string {

	pc, file, line, _ := runtime.Caller(frame)

	_, fileName := path.Split(file)

	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	packageName := parts[0]
	funcName := strings.Join(parts[1:], ".")

	return packageName + ": " + fileName + ": " + funcName + "-" + fmt.Sprintf("%d", line) + ": "

}

func Info(msg string) {
	l.infoLogger.Println(getCallerInfo(2) + msg)
}

func Warn(msg string) {
	l.warnLogger.Println(getCallerInfo(2) + msg)
}

func Error(msg string) {
	l.errorLogger.Println(getCallerInfo(2) + msg)
}
