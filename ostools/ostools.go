package ostools

import (
	"io"
	"os"
)

type Service struct{}

// NewService provides a service for interacting with the OS, which can easily be mocked
// This service implements an interface which is what gets passed to other services
// Please see mocks.go in the same package for the mock implementation
func NewService() Service {
	return Service{}
}

// Open opens a file, it is isolated to its most basic form for easy mocking
func (s Service) Open(path string) (io.Reader, func() error, error) {
	f, err := os.Open(path)
	return f, f.Close, err
}
