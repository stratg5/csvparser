package csv

import (
	"empora/ostools"
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
			desc:             "only headers",
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
