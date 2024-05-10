package address

import (
	"csvparser/entities"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

type AddressProvider interface {
	BuildLookupsFromAddresses(addresses []entities.Address) []*street.Lookup
	BuildAddressesFromRawData(data [][]string) []entities.Address
	BuildRawDataFromLookups(addresses []entities.Address, lookups []*street.Lookup) ([]string, error)

	SendLookups(lookups ...*street.Lookup) error
}

type lookupSender interface {
	SendLookups(lookups ...*street.Lookup) error
}
