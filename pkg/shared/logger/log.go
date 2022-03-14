package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	f      *os.File
	logger *log.Logger

	DefaultCallerDepth = 2
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	logPrefix          = ""
	eFlag              = ""
	eFunc              = ""
	eFile              = ""
	eLine              = -1
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	now := time.Now()
	filePath := "audit_log/"
	fileName := fmt.Sprintf("%s.log", now.Format("20060102150405"))
	f, err = MustOpen(filePath, fileName)
	if err != nil {
		log.Fatalf("logging setup err: %v", err)
	}
	logger = log.New(f, "", log.LstdFlags)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	log.Println(v...)
	logger.Println(v...)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	log.Println(v...)
	logger.Println(v...)
}
func Warning(v ...interface{}) {
	setPrefix(WARNING)
	log.Println(v...)
	logger.Println(v...)
}

func setPrefix(level Level) {
	function, file, line, ok := runtime.Caller(DefaultCallerDepth)

	if ok {
		s := strings.Split(runtime.FuncForPC(function).Name(), ".")
		_, fn := s[0], s[1]
		logPrefix = fmt.Sprintf("[%s][SYS][%s][%s:%d]", levelFlags[level], fn, filepath.Base(file), line)
		eFlag = levelFlags[level]
		eFunc = fn
		eFile = filepath.Base(file)
		eLine = line
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
