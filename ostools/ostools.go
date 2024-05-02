package ostools

import (
	"io"
	"os"
)

type Service struct{}

// NewService provides a service for interacting with the OS, which can easily be mocked
func NewService() Service {
	return Service{}
}

// Open opens a file, it is isolated to its most basic form for easy mocking
func (s Service) Open(path string) (io.Reader, func() error, error) {
	f, err := os.Open(path)
	return f, f.Close, err
}
