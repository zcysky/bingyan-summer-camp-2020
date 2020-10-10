package util

import (
	"io"
	"log"
	"os"
)

var LogWriter io.Writer

func initLog() {
	logFile, err := os.OpenFile("qq-bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("util: error when init logger")
		log.Panic(err)
	}

	LogWriter = io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(LogWriter)
}
