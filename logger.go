package logger

import (
	"log"
)

type Logger struct {
	Writer
	ID        string
	Formatter Formatter
	Prefixs   []Prefix
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
func (l *Logger) Log(logdata Log) {
	var output string
	if l.Formatter == nil {
		output = DefaultFormatter.Format(logdata)
	} else {
		output = l.Formatter.Format(logdata)
	}
	l.LogString(output)
}

func (l *Logger) LogString(s string) {
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
func (l *Logger) SetFormatter(f Formatter) *Logger {
	l.Formatter = f
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
		ID:        l.ID,
		Writer:    l.Writer,
		Formatter: l.Formatter,
		Prefixs:   p,
	}
}

func (l *Logger) SubLogger() *Logger {
	logger := l.Clone()
	logger.Writer = l
	return logger
}

func NewLogger() *Logger {
	return &Logger{
		Writer:  Null,
		Prefixs: []Prefix{},
	}
}
func createLogger(w Writer, id string, f Formatter, p ...Prefix) *Logger {
	return &Logger{
		ID:        id,
		Writer:    w,
		Formatter: f,
		Prefixs:   p,
	}
}
