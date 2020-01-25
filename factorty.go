package logger

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"sync"
)

type Option interface {
	ApplyTo(*Logger) error
}

type Factory func(u *url.URL) (Option, error)

func InitLogger(l *Logger, option string) error {
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
	o, err := factoryi(u)
	if err != nil {
		return err
	}
	return o.ApplyTo(l)
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
