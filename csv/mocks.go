package csv

type MockIOWriter struct {
	WriteFn func(p []byte) (n int, err error)
}

func (m MockIOWriter) Write(p []byte) (n int, err error) {
	return m.WriteFn(p)
}