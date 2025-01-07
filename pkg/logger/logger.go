package logger

import (
	"log"
	"os"
)

func InitLogger() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
