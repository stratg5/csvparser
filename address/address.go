package address

import (
	"empora/entities"
	"fmt"
	"strings"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

type Service struct {
	LookupClient LookupSender
}

// NewService generates a new address service
// This client currently has the ability to perform lookups but could be extended in the future
func NewService(lookupSender LookupSender) Service {
	return Service{
		LookupClient: lookupSender,
	}
}

func (s Service) BuildLookups(addresses []entities.Address) []*street.Lookup {
	lookups := []*street.Lookup{}
	for _, address := range addresses {
		lookups = append(lookups, &street.Lookup{
			Street:  address.Street,
			City:    address.Street,
			ZIPCode: address.ZipCode,
		})
	}

	return lookups
}

func (s Service) SendLookups(lookups ...*street.Lookup) error {
	err := s.LookupClient.SendLookups(lookups...)
	if err != nil {
		return fmt.Errorf("error while sending lookups: %w", err)
	}

	return nil
}

// BuildAddresses takes in the raw CSV data and builds an entity array
func (s Service) BuildAddressesFromRawData(data [][]string) []entities.Address {
	addresses := []entities.Address{}
	for _, row := range data {
		// TODO check if the initial rows are correct, just skip index 0?
		// check if the row is an invalid length
		if len(row) < 3 || len(row) > 3 {
			originString := ""
			for _, col := range row {
				originString += col
			}

			address := entities.Address{
				OriginString: originString,
				Valid:        false,
			}

			addresses = append(addresses, address)
			continue
		}

		// trim the leading and trailing spaces from the data
		address := entities.Address{
			Street:  strings.TrimSpace(row[0]),
			City:    strings.TrimSpace(row[1]),
			ZipCode: strings.TrimSpace(row[2]),
			Valid:   true,
		}

		addresses = append(addresses, address)
	}
	return addresses
}
