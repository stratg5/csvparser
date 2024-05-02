package address

import (
	"empora/entities"
	"errors"
	"reflect"
	"testing"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

func TestBuildLookupsFromAddresses(t *testing.T) {
	type test struct {
		desc             string
		addresses         []entities.Address
		expectedLookups []*street.Lookup
	}

	tests := []test{
		{
			desc:             "empty address array gives empty lookups",
			addresses: []entities.Address{},
			expectedLookups: []*street.Lookup{},
		},
		{
			desc:             "with one address",
			addresses: []entities.Address{
				{
					Street: "street1",
					City: "city1",
					ZipCode: "zip1",
				},
			},
			expectedLookups: []*street.Lookup{
					{
						Street: "street1",
						City: "city1",
						ZIPCode: "zip1",
					},
			},
		},
		{
			desc:             "with multiple addresses",
			addresses: []entities.Address{
				{
					Street: "street1",
					City: "city1",
					ZipCode: "zip1",
				},
				{
					Street: "street2",
					City: "city2",
					ZipCode: "zip2",
				},
			},
			expectedLookups: []*street.Lookup{
					{
						Street: "street2",
						City: "city2",
						ZIPCode: "zip2",
					},
			},
		},
	}

	for _, test := range tests {
		s := Service{}
		lookups := s.BuildLookupsFromAddresses(test.addresses)

		for _, address := range test.addresses {
			found := false
			for _, lookup := range lookups {
				if address.Street == lookup.Street && address.City == lookup.City && address.ZipCode == lookup.ZIPCode {
					found = true
				}
			}
			if !found {
				t.Fatalf("could not find address in lookup array")
				t.Fail()
			}
		}
	}
}

func TestBuildRawDataFromLookups(t *testing.T) {
	type test struct {
		desc             string
		addresses         []entities.Address
		lookups []*street.Lookup
		expectedContents []string
		expectedErr string
	}

	tests := []test{
		{
			desc:             "empty addresses and lookups",
			expectedContents: []string{},
		},
		{
			desc:             "addresses only",
			addresses: []entities.Address{
				{
					Street: "street1",
					City: "city1",
					ZipCode: "zip1",
				},
			},
			expectedContents: []string{},
			expectedErr: "address and lookup lengths don't match",
		},
		{
			desc:             "lookups only",
			lookups: []*street.Lookup{
				{
					Street: "street1",
					City: "city1",
					ZIPCode: "zip1",
				},
			},
			expectedContents: []string{},
			expectedErr: "address and lookup lengths don't match",
		},
		{
			desc:             "with invalid address and lookup",
			addresses: []entities.Address{
				{
					Street: "street1",
					City: "city1",
					OriginString: "street1, city1",
					Valid: false,
				},
			},
			lookups: []*street.Lookup{
				{
					Street: "street1",
					City: "city1",
					ZIPCode: "zip1",
				},
			},
			expectedContents: []string{"street1, city1 -> Invalid Address"},
		},
		{
			desc:             "with valid address and lookup",
			addresses: []entities.Address{
				{
					Street: "street1",
					City: "city1",
					ZipCode: "12345",
					Valid: true,
				},
			},
			lookups: []*street.Lookup{
				{
					Street: "street1",
					City: "city1",
					ZIPCode: "zip1",
					Results: []*street.Candidate{
						{
							DeliveryLine1: "street1",
							LastLine: "city1, zip1",
						},
					},
				},
			},
			expectedContents: []string{"street1, city1, 12345 -> street1, city1, zip1"},
		},
	}

	for _, test := range tests {
		s := Service{}
		data, err := s.BuildRawDataFromLookups(test.addresses, test.lookups)
		if err != nil {
			if test.expectedErr == "" {
				t.Fatalf(test.desc + ": expected no error but got one: %s", err.Error())
			}

			if err.Error() != test.expectedErr {
				t.Fatalf(test.desc + ": expected no error but got one: %s", err.Error())
			}

			continue
		}

		if !reflect.DeepEqual(data, test.expectedContents) {
			t.Fatalf(test.desc+" failed equality check", data, test.expectedContents)
		}
	}
}

func TestBuildAddressesFromRawData(t *testing.T) {
	type test struct {
		desc             string
		data         [][]string
		expectedContents []entities.Address
	}

	tests := []test{
		{
			desc:             "empty data",
			expectedContents: []entities.Address{},
		},
		{
			desc:             "just headers",
			data: [][]string{
				{
					"Street", "City", "Zip",
				},
			},
			expectedContents: []entities.Address{},
		},
		{
			desc:             "headers plus data",
			data: [][]string{
				{
					"Street", "City", "Zip",
				},
				{
					"123 Main Street", "Columbus", "43212",
				},
			},
			expectedContents: []entities.Address{
				{
					Street: "123 Main Street",
					City: "Columbus",
					ZipCode: "43212",
					Valid: true,
				},
			},
		},
		{
			desc:             "invalid row length",
			data: [][]string{
				{
					"Street", "City", "Zip",
				},
				{
					"123 Main Street", "Columbus",
				},
			},
			expectedContents: []entities.Address{
				{
					OriginString: "123 Main Street, Columbus",
					Valid: false,
				},
			},
		},
	}

	for _, test := range tests {
		s := Service{}
		addresses := s.BuildAddressesFromRawData(test.data)

		if !reflect.DeepEqual(addresses, test.expectedContents) {
			t.Fatalf(test.desc+" failed equality check", addresses, test.expectedContents)
		}
	}
}

func TestSendLookups(t *testing.T) {
	type test struct {
		desc             string
		data []*street.Lookup
		lookupErr error
		expectedErr string
	}

	tests := []test{
		{
			desc:             "empty data",
		},
		{
			desc:             "with data, no error",
			data: []*street.Lookup{
				{
					Street: "street",
				},
			},
		},
		{
			desc:             "with data and error",
			data: []*street.Lookup{
				{
					Street: "street",
				},
			},
			lookupErr: errors.New("something went wrong"),
			expectedErr: "error while sending lookups: something went wrong",
		},
	}

	for _, test := range tests {
		s := Service{
			LookupClient: MockLookupSender{
				SendLookupsFn: func(lookups ...*street.Lookup) error {
					return test.lookupErr
				},
			},
		}
		err := s.SendLookups(test.data...)
		if err != nil {
			if test.expectedErr == "" {
				t.Fatalf(test.desc + ": expected no error but got one: %s", err.Error())
			}

			if err.Error() != test.expectedErr {
				t.Fatalf(test.desc + ": expected no error but got one: %s", err.Error())
			}

			continue
		}
	}
}