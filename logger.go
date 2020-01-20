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

func (l *Logger) logBytes(p []byte) error {
	var err error
	var prefix = []byte(getPrefixs(l, DefaultPrefixSep, l.Prefixs...))
	_, err = l.Writer.Write(prefix)
	if err != nil {
		return err
	}
	_, err = l.Writer.Write(p)
	return err
}
func (l *Logger) Log(v ...interface{}) {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	var data []byte

	if l.Formatter == nil {
		data, err = DefaultFormatter.Format(v...)
	} else {
		data, err = l.Formatter.Format(v...)
	}
	if err != nil {
		return
	}
	err = l.logBytes(data)
	return
}

func (l *Logger) LogBytes(p []byte) {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	err = l.logBytes(p)
	return
}

func (l *Logger) LogString(s string) {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	err = l.logBytes([]byte(s))
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
	return &Logger{}
}
func createLogger(w Writer, id string, f Formatter, p ...Prefix) *Logger {
	return &Logger{
		ID:        id,
		Writer:    w,
		Formatter: f,
		Prefixs:   p,
	}
}
