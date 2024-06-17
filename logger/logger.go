package logger

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func InitLogger() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
	log.Println("INFO: " + message)
}

func LogError(message string, err error) {
	log.Printf("ERROR: %s - %s\n", message, err)
	sentry.CaptureException(err)
}

func LogFatal(message string, err error) {
	log.Fatalf("FATAL: %s - %s\n", message, err)
	sentry.CaptureException(err)
}
