package address

import (
	"empora/entities"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

type AddressProvider interface {
	BuildLookups(addresses []entities.Address) []*street.Lookup
	SendLookups(lookups ...*street.Lookup) error
	BuildAddressesFromRawData(data [][]string) []entities.Address
}

type LookupSender interface {
	SendLookups(lookups ...*street.Lookup) error
}
