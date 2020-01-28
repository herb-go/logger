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

var DefaultFormatter Formatter = SeparatedFormatter(" ")
