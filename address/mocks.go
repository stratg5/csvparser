package address

import (
	"empora/entities"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

type MockLookupSender struct {
	SendLookupsFn func(lookups ...*street.Lookup) error
}

func (m MockLookupSender) SendLookups(lookups ...*street.Lookup) error {
	return m.SendLookupsFn(lookups...)
}

type MockAddressProvider struct {
	BuildLookupsFromAddressesFn func(addresses []entities.Address) []*street.Lookup
	BuildAddressesFromRawDataFn func(data [][]string) []entities.Address
	BuildRawDataFromLookupsFn func(addresses []entities.Address, lookups []*street.Lookup) ([]string, error)

	SendLookupsFn func(lookups ...*street.Lookup) error
}

func (m MockAddressProvider) BuildLookupsFromAddresses(addresses []entities.Address) []*street.Lookup {
	return m.BuildLookupsFromAddressesFn(addresses)
}

func (m MockAddressProvider) BuildAddressesFromRawData(data [][]string) []entities.Address {
 return m.BuildAddressesFromRawDataFn(data)
}

func (m MockAddressProvider) BuildRawDataFromLookups(addresses []entities.Address, lookups []*street.Lookup) ([]string, error) {
return m.BuildRawDataFromLookupsFn(addresses, lookups)
}

func (m MockAddressProvider) SendLookups(lookups ...*street.Lookup) error {
return m.SendLookupsFn(lookups...)
}