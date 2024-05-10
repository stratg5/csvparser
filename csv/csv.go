package csv

import (
	"csvparser/ostools"
	"encoding/csv"
	"errors"
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
// Keeping the parsing logic separate from the reading of the csv in case this needs extended
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

func (s Service) WriteCSV(path string, records []string) error {
	csvFile, close, err := s.ostool.Create(path)
	if err != nil {
		return fmt.Errorf("error creating csv file: %w", err)
	}
	defer close()


	for _, record := range records {
		_, writeErr := csvFile.Write([]byte(record + "\n"))
		if err != nil {
			err = errors.Join(err, writeErr)
		}
	}

	return err
}

