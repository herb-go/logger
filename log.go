package logger

type Log interface {
	DefaultLogFields() []interface{}
	LogFields() map[string]interface{}
}

type LogFields []string

func (f *LogFields) NewLog() *PlainLog {
	l := NewPlainLog()
	l.defaultLogFieldNames = *f
	return l
}

func NewLogFields(fields ...string) *LogFields {
	f := LogFields(fields)
	return &f
}

type PlainLog struct {
	defaultLogFieldNames []string
	logFields            map[string]interface{}
}

func NewPlainLog() *PlainLog {
	return &PlainLog{
		logFields: map[string]interface{}{},
	}
}

func (l *PlainLog) DefaultLogFields() []interface{} {
	f := make([]interface{}, len(l.defaultLogFieldNames))
	for k, v := range l.defaultLogFieldNames {
		f[k] = l.logFields[v]
	}
	return f
}
func (l *PlainLog) LogFields() map[string]interface{} {
	return l.logFields
}

func (l *PlainLog) WithField(name string, value interface{}) *PlainLog {
	l.logFields[name] = value
	return l
}
