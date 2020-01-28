package logger

import (
	"io"
	"io/ioutil"
	"os"
	"sync"
)

var newline = []byte{'\n'}

type Writer interface {
	Open() error
	Close() error
	Reopen() error
	WriteLine(s string) error
}

type IOWriter struct {
	Writer io.Writer
}

func (o *IOWriter) WriteLine(s string) error {
	var err error
	_, err = o.Writer.Write([]byte(s))
	if err != nil {
		return err
	}
	_, err = o.Writer.Write(newline)
	return err
}

func (o *IOWriter) Open() error {
	return nil
}
func (o *IOWriter) Close() error {
	return nil
}

func (o *IOWriter) Reopen() error {
	return nil
}

var Stdout Writer = &IOWriter{
	Writer: os.Stdout,
}

var Stderr Writer = &IOWriter{
	Writer: os.Stderr,
}

var Null Writer = &IOWriter{
	Writer: ioutil.Discard,
}

type FileWriter struct {
	lock sync.RWMutex
	Path string
	Mode os.FileMode
	file *os.File
}

func (o *FileWriter) Open() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	file, err := os.OpenFile(o.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, o.Mode)
	if err != nil {
		return err
	}
	o.file = file
	return nil
}
func (o *FileWriter) Close() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.file = nil
	return o.file.Close()
}

func (o *FileWriter) WriteLine(s string) error {
	o.lock.RLock()
	defer o.lock.RUnlock()
	var err error
	_, err = o.file.Write([]byte(s))
	if err != nil {
		return err
	}
	_, err = o.file.Write(newline)
	return err
}
func (o *FileWriter) Reopen() error {
	var err error
	o.lock.Lock()
	defer o.lock.Unlock()
	err = o.file.Close()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(o.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, o.Mode)
	if err != nil {
		return err
	}
	o.file = file
	return nil
}
func NewFileWriter(path string, mode os.FileMode) *FileWriter {
	return &FileWriter{
		Path: path,
		Mode: mode,
	}
}
