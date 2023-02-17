package pkg

import (
	"log"
	"os"
)

func WarningLogger() *log.Logger {
	return log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
}

func InfoLogger() *log.Logger {
	return log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
}

func ErrorLogger() *log.Logger {
	return log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime)
}
