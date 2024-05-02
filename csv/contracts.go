package csv

type CSVProvider interface {
	ReadCSV(path string) ([][]string, error)
}
