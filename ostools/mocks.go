package ostools

import (
	"io"
)

type Mock struct {
	OpenFn   func(path string) (io.Reader, func() error, error)
	CreateFn func(path string) (io.Writer, func() error, error)
}

func (m Mock) Open(path string) (io.Reader, func() error, error) {
	return m.OpenFn(path)
}

func (m Mock) Create(path string) (io.Writer, func() error, error) {
	return m.CreateFn(path)
}
