package ostools

import "io"

type OSTooler interface {
	Open(path string) (io.Reader, func() error, error)
}
