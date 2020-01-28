package logger

type Log interface {
	DefaultLogFields() []interface{}
	LogFields() map[string]interface{}
}
