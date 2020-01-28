package logger

import (
	"log"
)

var PanicLogger *Logger
var FatalLogger *Logger
var ErrorLogger *Logger
var WarningLogger *Logger
var InfoLogger *Logger
var PrintLogger *Logger
var TraceLogger *Logger
var DebugLogger *Logger

func ResetBuiltinLoggers() {
	PanicLogger = createLogger(Stderr, "Panic", DefaultTimePrefix, PrefixID)
	FatalLogger = createLogger(Stderr, "Fatal", DefaultTimePrefix, PrefixID)
	ErrorLogger = createLogger(Stderr, "Error", DefaultTimePrefix, PrefixID)
	WarningLogger = createLogger(Stdout, "Warning", DefaultTimePrefix, PrefixID)
	InfoLogger = createLogger(Stdout, "Info", DefaultTimePrefix, PrefixID)
	PrintLogger = createLogger(Stdout, "Print")
	TraceLogger = createLogger(Null, "Trace", DefaultTimePrefix, PrefixID)
	DebugLogger = createLogger(Null, "Debug", DefaultTimePrefix, PrefixID)
}

func Reopen(w ...Writer) {
	var err error
	for k := range w {
		err = w[k].Reopen()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReopenBuiltinLoggers() {
	Reopen(PanicLogger, FatalLogger, ErrorLogger, PrintLogger, WarningLogger, InfoLogger, TraceLogger, DebugLogger)
}
func Panic(s string) {
	PanicLogger.Log(s)
}
func Fatal(s string) {
	FatalLogger.Log(s)
}
func Error(s string) {
	ErrorLogger.Log(s)
}
func Warning(s string) {
	WarningLogger.Log(s)
}
func Info(s string) {
	InfoLogger.Log(s)
}
func Print(s string) {
	PrintLogger.Log(s)
}
func Trace(s string) {
	TraceLogger.Log(s)
}
func Debug(s string) {
	DebugLogger.Log(s)
}

func EnableDevelopmengLoggers() {
	TraceLogger.Writer = Stdout
	DebugLogger.Writer = Stdout
}

func init() {
	ResetBuiltinLoggers()
}
