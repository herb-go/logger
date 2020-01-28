package logger

import (
	"fmt"
	"strings"
)

type Formatter interface {
	Format(Log) string
}

type SeparatedFormatter string

func (f SeparatedFormatter) Format(l Log) string {
	v := l.DefaultLogFields()
	var data = make([]string, len(v))
	for i := range v {
		data[i] = fmt.Sprint(v[i])
	}
	return strings.Join(data, string(f))
}

var SpaceSeparatedFormatter = SeparatedFormatter(" ")
var DefaultFormatter Formatter = SpaceSeparatedFormatter

type ReplacementFormater string

func (f ReplacementFormater) Format(l Log) string {
	if f == "" {
		return SpaceSeparatedFormatter.Format(l)
	}
	replacement := []string{}
	fields := l.LogFields()
	for k := range fields {
		replacement = append(replacement, "{{"+k+"}}", fmt.Sprint(fields[k]))
	}
	replacement = append(replacement, "\\{", "{", "\\}", "}", "\\\\", "\\")
	return strings.NewReplacer(replacement...).Replace(string(f))
}

type FormatLogger struct {
	Logger    *Logger
	Formatter Formatter
}

func (l *FormatLogger) FormatAndLog(log Log) {
	if l.Formatter != nil {
		l.Logger.Log(l.Formatter.Format(log))

	} else {
		l.Logger.Log(DefaultFormatter.Format(log))
	}
}
