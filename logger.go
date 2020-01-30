package logger

import (
	"log"
)

type Logger struct {
	Writer
	ID      string
	Prefixs []Prefix
}

func (l *Logger) ReplaceWriter(w Writer) error {
	if l.Writer != nil {
		err := l.Writer.Close()
		if err != nil {
			return err
		}
	}
	l.Writer = w
	return l.Writer.Open()
}

func (l *Logger) Log(s string) {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	output := getPrefixs(l, DefaultPrefixSep, l.Prefixs...)
	output = output + s
	err = l.Writer.WriteLine(output)
	return
}

func (l *Logger) SetWriter(w Writer) *Logger {
	l.Writer = w
	return l
}
func (l *Logger) SetID(id string) *Logger {
	l.ID = id
	return l
}

func (l *Logger) SetPrefixs(p ...Prefix) *Logger {
	l.Prefixs = p
	return l
}
func (l *Logger) AppendPrefixs(p ...Prefix) *Logger {
	l.Prefixs = append(l.Prefixs, p...)
	return l
}
func (l *Logger) Clone() *Logger {
	p := make([]Prefix, len(l.Prefixs))
	copy(p, l.Prefixs)
	return &Logger{
		ID:      l.ID,
		Writer:  l.Writer,
		Prefixs: p,
	}
}

func (l *Logger) SubLogger() *Logger {
	logger := l.Clone()
	logger.Writer = l
	return logger
}
func (l *Logger) ForamtLogger() *FormatLogger {
	return &FormatLogger{
		Logger:    l,
		Formatter: nil,
	}
}
func NewLogger() *Logger {
	return &Logger{
		Writer:  Null,
		Prefixs: []Prefix{},
	}
}
func createLogger(w Writer, id string, p ...Prefix) *Logger {
	return &Logger{
		ID:      id,
		Writer:  w,
		Prefixs: p,
	}
}
