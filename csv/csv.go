package csv

import (
	"empora/ostools"
	"encoding/csv"
	"fmt"
)

type Service struct {
	ostool ostools.OSTooler
}

func NewService(ostool ostools.OSTooler) Service {
	return Service{
		ostool: ostool,
	}
}

// ReadCSV reads the csv contents at the given path and returns the raw data plus error
func (s Service) ReadCSV(path string) ([][]string, error) {
	// Open the CSV file
	file, close, err := s.ostool.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error while opening the csv file: %w", err)
	}
	defer close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}
