package csv

type MockIOWriter struct {
	WriteFn func(p []byte) (n int, err error)
}

func (m MockIOWriter) Write(p []byte) (n int, err error) {
	return m.WriteFn(p)
}

type MockCSVProvider struct {
	ReadCSVFn  func(path string) ([][]string, error)
	WriteCSVFn func(path string, records []string) error
}

func (m MockCSVProvider) ReadCSV(path string) ([][]string, error) {
	return m.ReadCSVFn(path)
}

func (m MockCSVProvider) WriteCSV(path string, records []string) error {
	return m.WriteCSVFn(path, records)
}
