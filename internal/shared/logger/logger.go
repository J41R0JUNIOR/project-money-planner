package logger

import "log"

func Info(msg string, args ...any) {
	log.Printf(msg, args...)
}
