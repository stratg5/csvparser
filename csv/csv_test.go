package csv

import (
	"csvparser/ostools"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadCSV(t *testing.T) {
	type test struct {
		desc             string
		contents         string
		expectedContents [][]string
	}

	tests := []test{
		{
			desc:             "should read the data into the expected format",
			contents:         `Street,City,Zip Code`,
			expectedContents: [][]string{{"Street", "City", "Zip Code"}},
		},
	}

	for _, test := range tests {
		s := Service{
			ostool: ostools.Mock{
				OpenFn: func(path string) (io.Reader, func() error, error) {
					return strings.NewReader(test.contents), func() error { return nil }, nil
				},
			},
		}
		data, err := s.ReadCSV("test.csv")
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(data, test.expectedContents) {
			t.Fatalf(test.desc+" failed equality check", data, test.expectedContents)
		}

	}
}

func TestWriteCSV(t *testing.T) {
	type test struct {
		desc             string
		records []string
		expectedContents string
	}

	tests := []test{
		{
			desc:             "should write data with a newline",
			records:        []string{"hello"},
			expectedContents: "hello\n",
		},
	}

	for _, test := range tests {
		data := ""
		s := Service{
			ostool: ostools.Mock{
				CreateFn: func(path string) (io.Writer, func() error, error) {
					return MockIOWriter{
						WriteFn: func(p []byte) (n int, err error) {
							data = string(p)
							return len(p), err
						},
					}, func() error { return nil }, nil
				},
			},
		}

		err := s.WriteCSV("test.csv", test.records)
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(data, test.expectedContents) {
			t.Fatalf(test.desc+" failed equality check", data, test.expectedContents)
		}
	}
} 
