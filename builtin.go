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
	PanicLogger = createLogger(Stderr, "Panic", nil, DefaultTimePrefix, PrefixID)
	FatalLogger = createLogger(Stderr, "Fatal", nil, DefaultTimePrefix, PrefixID)
	ErrorLogger = createLogger(Stderr, "Error", nil, DefaultTimePrefix, PrefixID)
	WarningLogger = createLogger(Stdout, "Warning", nil, DefaultTimePrefix, PrefixID)
	InfoLogger = createLogger(Stdout, "Info", nil, DefaultTimePrefix, PrefixID)
	PrintLogger = createLogger(Stdout, "Print", nil)
	TraceLogger = createLogger(Null, "Trace", nil, DefaultTimePrefix, PrefixID)
	DebugLogger = createLogger(Null, "Debug", nil, DefaultTimePrefix, PrefixID)
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
	PanicLogger.LogString(s)
}
func Fatal(s string) {
	FatalLogger.LogString(s)
}
func Error(s string) {
	ErrorLogger.LogString(s)
}
func Warning(s string) {
	WarningLogger.LogString(s)
}
func Info(s string) {
	InfoLogger.LogString(s)
}
func Print(s string) {
	PrintLogger.LogString(s)
}
func Trace(s string) {
	TraceLogger.LogString(s)
}
func Debug(s string) {
	DebugLogger.LogString(s)
}

func EnableDevelopmengLoggers() {
	TraceLogger.Writer = Stdout
	DebugLogger.Writer = Stdout
}

func init() {
	ResetBuiltinLoggers()
}
