package ostools

import "io"

type OSTooler interface {
	Open(path string) (io.Reader, func() error, error)
	Create(path string) (io.Writer, func() error, error)
}
