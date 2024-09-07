package logger

import (
	"fmt"
	"integriWatch/utils"
	"log"
	"os"
	"time"
)

const (
	logFileDir    = "Logs"
	logFilePrefix = "IntegriShield-log-"
)

var (
	terminalLogger *log.Logger
	fileLogger     *log.Logger
	logFileName    string
	logFile        *os.File
	err            error
	fileInit       bool
)

func Init(appDataDir string) {

	// Ensure Logs dir
	utils.EnsureDir(
		fmt.Sprintf(
			"%s%s%s",
			appDataDir,
			string(os.PathSeparator),
			logFileDir,
		),
	)

	// Create the terminal logger
	terminalLogger = log.New(os.Stdout, "IntegriWatch:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create file logger
	logFileName = fmt.Sprintf(
		"%s%s%s%s%s%d.log",
		appDataDir,
		string(os.PathSeparator),
		logFileDir,
		string(os.PathSeparator),
		logFilePrefix,
		time.Now().Unix(),
	)
	logFile, err = os.Create(logFileName)
	fileInit = true
	if err != nil {
		fileInit = false
		Log(fmt.Sprintf("failed to create log file: %v", err))
	}
	fileLogger = log.New(logFile, "IntegriWatch:\t", log.Ldate|log.Ltime)

	// Log configuration update
	Log("logger initialised\n")
}

func Log[T any](log T) {
	terminalLogger.Print(log)
	if fileInit {
		fileLogger.Print(log)
	}
}

func CleanUp() {
	if err = logFile.Close(); err != nil {
		Log("log file already closed\n")
	}
}
