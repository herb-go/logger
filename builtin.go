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

func reopen(w ...Writer) {
	var err error
	for k := range w {
		err = w[k].Reopen()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReopenBuiltinLoggers() {
	reopen(PanicLogger, FatalLogger, ErrorLogger, PrintLogger, WarningLogger, InfoLogger, TraceLogger, DebugLogger)
}
func Panic(v ...interface{}) {
	PanicLogger.Log(v...)
}
func Fatal(v ...interface{}) {
	FatalLogger.Log(v...)
}
func Error(v ...interface{}) {
	ErrorLogger.Log(v...)
}
func Warning(v ...interface{}) {
	WarningLogger.Log(v...)
}
func Info(v ...interface{}) {
	InfoLogger.Log(v...)
}
func Print(v ...interface{}) {
	PrintLogger.Log(v...)
}
func Trace(v ...interface{}) {
	TraceLogger.Log(v...)
}
func Debug(v ...interface{}) {
	DebugLogger.Log(v...)
}

func EnableDevelopmengLoggers() {
	TraceLogger.Writer = Stdout
	DebugLogger.Writer = Stdout
}

func init() {
	ResetBuiltinLoggers()
}
