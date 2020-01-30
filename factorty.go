package logger

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"sort"
	"sync"
)

type Option interface {
	ApplyTo(*Logger) error
}

type OptionFunc func(*Logger) error

func (o OptionFunc) ApplyTo(l *Logger) error {
	return o(l)
}

type Factory func(u *url.URL, loader func(v interface{}) error) (Option, error)

func InitLogger(l *Logger, option string, loader func(v interface{}) error) error {
	factorysMu.RLock()
	u, err := url.Parse(option)
	if err != nil {
		return err
	}
	name := u.Scheme
	factoryi, ok := factories[name]
	factorysMu.RUnlock()
	if !ok {
		return fmt.Errorf("logger: unknown driver %q (forgotten import?)", name)
	}
	o, err := factoryi(u, loader)
	if err != nil {
		return err
	}
	err = o.ApplyTo(l)
	if err != nil {
		return err
	}
	return nil
}

var (
	factorysMu sync.RWMutex
	factories  = make(map[string]Factory)
)

// Register makes a driver creator available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, f Factory) {
	factorysMu.Lock()
	defer factorysMu.Unlock()
	if f == nil {
		panic(errors.New("logger: Register logger factory is nil"))
	}
	if _, dup := factories[name]; dup {
		panic(errors.New("logger: Register called twice for factory " + name))
	}
	factories[name] = f
}

//UnregisterAll unregister all driver
func UnregisterAll() {
	factorysMu.Lock()
	defer factorysMu.Unlock()
	// For tests.
	factories = make(map[string]Factory)
}

//Factories returns a sorted list of the names of the registered factories.
func Factories() []string {
	factorysMu.RLock()
	defer factorysMu.RUnlock()
	var list []string
	for name := range factories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func RegisterBuiltinFactory() {
	Register("", func(u *url.URL, loader func(v interface{}) error) (Option, error) {
		return OptionFunc(func(l *Logger) error {
			var err error
			if u.Host != "" {
				return fmt.Errorf("logger:builtin logger formaterror %s", u.String())
			}
			switch u.Path {
			case "":
			case "stdout":
				err = l.ReplaceWriter(Stdout)
			case "stderr":
				err = l.ReplaceWriter(Stderr)
			case "null":
				err = l.ReplaceWriter(Null)
			default:
				return fmt.Errorf("logger:unkown builtin logger output %s", u.Host)
			}
			if err != nil {
				return err
			}
			return nil
		}), nil
	})
}

func RegisterAbsoluteFileFactory() {
	Register("file", func(u *url.URL, loader func(v interface{}) error) (Option, error) {

		return OptionFunc(func(l *Logger) error {
			var err error
			err = l.ReplaceWriter(NewFileWriter(filepath.Join(u.Host, u.Path), 0660))
			if err != nil {
				return err
			}
			return nil
		}), nil
	})
}

func init() {
	RegisterBuiltinFactory()
	RegisterAbsoluteFileFactory()
}
