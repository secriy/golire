package util

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	LevelFatal LogLevel = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

var (
	level LogLevel = 1

	file *os.File
)

func Level() LogLevel {
	return level
}

func SetLevel(ll LogLevel) {
	level = ll

	var err error
	file, err = os.Create(fmt.Sprintf("golire_%s.log", time.Now().Format("20060102-15-04-05")))
	if err != nil {
		Fatal(err.Error())
	}
}

func SetLevelString(ll string) {
	switch ll {
	case "fatal":
		SetLevel(LevelFatal)
	case "error":
		SetLevel(LevelError)
	case "warning":
		SetLevel(LevelWarning)
	case "info":
		SetLevel(LevelInfo)
	case "debug":
		SetLevel(LevelDebug)
	default:
		SetLevel(LevelError)
	}
}

// writeLogFile write log text to file
func writeLogFile(text string) {
	if _, err := file.WriteString(text); err != nil {
		Error(err.Error())
	}
}

func logPrint(ll LogLevel, text string) {
	if ll == -1 || ll <= level {
		writeLogFile(text + "\n")

		fmt.Println(text)
	}
}

func Print(mod, text string) {
	logPrint(LogLevel(-1), fmt.Sprintf("[%s] [%s] %s", timeNow(), mod, text))
}

func Fatal(text string) {
	logPrint(LevelFatal, fmt.Sprintf("[%s] [%s] %s", timeNow(), "FATAL", text))
	os.Exit(1)
}

func Error(text string) {
	logPrint(LevelError, fmt.Sprintf("[%s] [%s] %s", timeNow(), "ERROR", text))
}

func Warning(text string) {
	logPrint(LevelWarning, fmt.Sprintf("[%s] [%s] %s", timeNow(), "WARNING", text))
}

func Info(text string) {
	logPrint(LevelInfo, fmt.Sprintf("[%s] [%s] %s", timeNow(), "INFO", text))
}

func Debug(text string) {
	logPrint(LevelDebug, fmt.Sprintf("[%s] [%s] %s", timeNow(), "DEBUG", text))
}

func timeNow() string {
	return time.Now().Format("15:04:05")
}
