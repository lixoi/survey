package logger

import (
	"fmt"
)

const (
	levelInfo  = "INFO"
	levelDebug = "DEBUG"
	levelWarn  = "WARNING"
	levelError = "ERROR"
	//levelErrDateBusy = "ErrDateBusy"
)

type Logger struct {
	level string
}

func New(level string) *Logger {
	return &Logger{level: level}
}

func (l Logger) Info(msg string) {
	if l.level == levelInfo {
		fmt.Println(levelInfo + ": " + msg)
	}
}

func (l Logger) Debug(msg string) {
	if l.level == levelDebug {
		fmt.Println(levelError + ": " + msg)
	}
}

func (l Logger) Wirning(msg string) {
	//if l.level == levelWarn {
	fmt.Println(levelError + ": " + msg)
	//}
}

func (l Logger) Error(msg string) {
	//if l.level == levelError {
	fmt.Println(levelError + ": " + msg)
	//}
}

// TODO
