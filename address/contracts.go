package address

import (
	"empora/entities"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

type AddressProvider interface {
	BuildLookupsFromAddresses(addresses []entities.Address) []*street.Lookup
	BuildAddressesFromRawData(data [][]string) []entities.Address
	BuildRawDataFromLookups(addresses []*street.Lookup) [][]string

	SendLookups(lookups ...*street.Lookup) error
}

type lookupSender interface {
	SendLookups(lookups ...*street.Lookup) error
}
