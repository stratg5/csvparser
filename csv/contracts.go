package csv

type CSVProvider interface {
	ReadCSV(path string) ([][]string, error)
	WriteCSV(path string, records []string) error
}
