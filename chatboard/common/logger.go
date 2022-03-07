package common

import (
	"log"
	"os"
)

type LogOutputMode int32

var logger *log.Logger

const (
	LogOutputScreen LogOutputMode = iota
	LogOutputFile   LogOutputMode = iota
)

///////////////////////////////////////
// should use zap ??

const (
	LogInfoPrefix    string = "[INFO] "
	LogWarningPrefix string = "[WARNING] "
	LogErrorPrefix   string = "[ERROR] "
)

func SetupLogger(mode LogOutputMode) (err error) {
	if mode == LogOutputFile {
		file, err := os.OpenFile("chatboard.log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("failed to open log file", err.Error())
		}
		logger = log.New(file, LogInfoPrefix,
			log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		logger = log.Default()
		logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	}
	return
}

func LogInfo() *log.Logger {
	logger.SetPrefix(LogInfoPrefix)
	return logger
}

func LogWarning() *log.Logger {
	logger.SetPrefix(LogWarningPrefix)
	return logger
}

func LogError() *log.Logger {
	logger.SetPrefix(LogErrorPrefix)
	return logger
}
