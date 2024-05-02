package ostools

import (
	"io"
)

type Mock struct {
	OpenFn func(path string) (io.Reader, func() error, error)
}

func (m Mock) Open(path string) (io.Reader, func() error, error) {
	return m.OpenFn(path)
}
