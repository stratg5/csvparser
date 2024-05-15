package driver

import (
	"csvparser/address"
	"csvparser/csv"
	"csvparser/entities"
	"errors"
	"testing"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

func TestParseCSVAndGenerateOutput(t *testing.T) {
	type test struct {
		desc string

		readCSVError      error
		writeCSVError     error
		sendLookupsError  error
		buildRawDataError error

		expectedWriteCSVCount       int
		expectedReadCSVCount        int
		expectedBuildAddressesCount int
		expectedBuildLookupsCount   int
		expectedSendLookupsCount    int
		expectedBuildRawDataCount   int
		expectedErr                 string
	}

	tests := []test{
		{
			desc:                        "no errors",
			expectedWriteCSVCount:       1,
			expectedReadCSVCount:        1,
			expectedBuildAddressesCount: 1,
			expectedBuildLookupsCount:   1,
			expectedSendLookupsCount:    1,
			expectedBuildRawDataCount:   1,
		},
		{
			desc: "read csv errors",

			readCSVError: errors.New("csv error"),

			expectedErr:                 "error while parsing csv: csv error",
			expectedWriteCSVCount:       0,
			expectedReadCSVCount:        1,
			expectedBuildAddressesCount: 0,
			expectedBuildLookupsCount:   0,
			expectedSendLookupsCount:    0,
			expectedBuildRawDataCount:   0,
		},
		{
			desc: "write csv errors",

			writeCSVError: errors.New("write csv error"),

			expectedErr:                 "error while writing output to csv: write csv error",
			expectedWriteCSVCount:       1,
			expectedReadCSVCount:        1,
			expectedBuildAddressesCount: 1,
			expectedBuildLookupsCount:   1,
			expectedSendLookupsCount:    1,
			expectedBuildRawDataCount:   1,
		},
		{
			desc: "send lookups errors",

			sendLookupsError: errors.New("send lookups error"),

			expectedErr:                 "error while sending lookups: send lookups error",
			expectedWriteCSVCount:       0,
			expectedReadCSVCount:        1,
			expectedBuildAddressesCount: 1,
			expectedBuildLookupsCount:   1,
			expectedSendLookupsCount:    1,
			expectedBuildRawDataCount:   0,
		},
		{
			desc: "build raw data errors",

			buildRawDataError: errors.New("build raw data error"),

			expectedErr:                 "error while building raw data: build raw data error",
			expectedWriteCSVCount:       0,
			expectedReadCSVCount:        1,
			expectedBuildAddressesCount: 1,
			expectedBuildLookupsCount:   1,
			expectedSendLookupsCount:    1,
			expectedBuildRawDataCount:   1,
		},
	}

	for _, test := range tests {
		readCSVCount := 0
		writeCSVCount := 0
		buildAddressesCount := 0
		buildLookupsCount := 0
		sendLookupsCount := 0
		buildRawDataCount := 0
		s := Service{
			csvSvc: csv.MockCSVProvider{
				ReadCSVFn: func(path string) ([][]string, error) {
					readCSVCount++
					return nil, test.readCSVError
				},
				WriteCSVFn: func(path string, records []string) error {
					writeCSVCount++
					return test.writeCSVError
				},
			},
			addressSvc: address.MockAddressProvider{
				BuildAddressesFromRawDataFn: func(data [][]string) []entities.Address {
					buildAddressesCount++
					return nil
				},
				BuildLookupsFromAddressesFn: func(addresses []entities.Address) []*street.Lookup {
					buildLookupsCount++
					return nil
				},
				SendLookupsFn: func(lookups ...*street.Lookup) error {
					sendLookupsCount++
					return test.sendLookupsError
				},
				BuildRawDataFromLookupsFn: func(addresses []entities.Address, lookups []*street.Lookup) ([]string, error) {
					buildRawDataCount++
					return nil, test.buildRawDataError
				},
			},
		}
		err := s.ParseCSVAndGenerateOutput("", "")
		if err != nil {
			if test.expectedErr == "" {
				t.Fatalf(test.desc+": expected no error but got one: %s", err.Error())
			}

			if err.Error() != test.expectedErr {
				t.Fatalf(test.desc+": expected no error but got one: %s", err.Error())
			}
		}

		validateCount(t, writeCSVCount, test.expectedWriteCSVCount, "write CSV")
		validateCount(t, readCSVCount, test.expectedReadCSVCount, "read CSV")
		validateCount(t, buildAddressesCount, test.expectedBuildAddressesCount, "build addresses")
		validateCount(t, buildLookupsCount, test.expectedBuildLookupsCount, "build lookups")
		validateCount(t, sendLookupsCount, test.expectedSendLookupsCount, "send lookups")
		validateCount(t, buildRawDataCount, test.expectedBuildRawDataCount, "build raw data")
	}
}

func validateCount(t *testing.T, count1, count2 int, name string) {
	if count1 != count2 {
		t.Fatalf("invalid count for: %s", name)
	}
}
