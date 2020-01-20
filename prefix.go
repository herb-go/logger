package logger

import (
	"time"
)

var DefaultPrefixSep = "|"

type Prefix interface {
	NewPrefix(*Logger) string
}

type TimePrefix struct {
	Layout string
}

func (p *TimePrefix) NewPrefix(*Logger) string {
	l := p.Layout
	if l == "" {
		l = DefaultTimeLayout
	}
	return time.Now().Format(l)
}

var DefaultTimePrefix = &TimePrefix{}

var DefaultTimeLayout = "2006-01-02 03:04:05"

type FixedPrefix string

func (p FixedPrefix) NewPrefix(*Logger) string {
	return string(p)
}

type PrefixFunc func(*Logger) string

func (f PrefixFunc) NewPrefix(l *Logger) string {
	return f(l)
}

var PrefixID = PrefixFunc(func(l *Logger) string {
	return l.ID
})

func getPrefixs(l *Logger, sep string, p ...Prefix) string {
	var data = ""
	for k := range p {
		data = data + p[k].NewPrefix(l) + sep
	}
	return data
}
