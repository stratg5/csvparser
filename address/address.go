package address

import (
	"empora/entities"
	"errors"
	"fmt"
	"strings"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

type Service struct {
	LookupClient lookupSender
}

// the address service handles any address related operations including sending and formatting
func NewService(lookupSender lookupSender) Service {
	return Service{
		LookupClient: lookupSender,
	}
}

func (s Service) BuildLookupsFromAddresses(addresses []entities.Address) []*street.Lookup {
	lookups := []*street.Lookup{}
	for _, address := range addresses {
		lookups = append(lookups, &street.Lookup{
			Street:  address.Street,
			City:    address.City,
			ZIPCode: address.ZipCode,
		})
	}

	return lookups
}

// BuildAddressesFromRawData takes in the raw CSV data and builds an address array
func (s Service) BuildAddressesFromRawData(data [][]string) []entities.Address {
	addresses := []entities.Address{}
	for idx, row := range data {
		// skip the column names
		if idx == 0 {
			continue
		}

		originString := ""
		for idx, col := range row {
			originString += col

			if idx < len(row) -1 && len(row) > 1 {
				originString += ", "
			}
		}

		// check if the row is an invalid length
		if len(row) < 3 || len(row) > 3 {
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
			OriginString: originString,
			Valid:   true,
		}

		addresses = append(addresses, address)
	}
	return addresses
}

// Builds the output based on address and lookup arrays
// If the array sizes don't match, return an error
func (s Service) BuildRawDataFromLookups(addresses []entities.Address, lookups []*street.Lookup) ([]string, error) {
	output := []string{}

	if len(addresses) != len(lookups) {
		return nil, errors.New("address and lookup lengths don't match")
	}

	for idx, lookup := range lookups {
		var addrLine string
		tempAddr := addresses[idx]

		if !tempAddr.Valid || len(lookup.Results) == 0 {
			
			addrLine = tempAddr.OriginString + " -> Invalid Address"
		} else {
			lookupAddr := lookup.Results[0]
			splitLastLine := strings.TrimSpace(lookupAddr.LastLine)
			addrParts := strings.Split(splitLastLine, " ")

			if len(addrParts) < 3 {
				addrLine = tempAddr.OriginString + " -> Invalid Address"
			} else {
				addrLine = tempAddr.Street+ ", " + tempAddr.City +", " + tempAddr.ZipCode + " -> " + lookupAddr.DeliveryLine1 + ", " + addrParts[0] + ", " + addrParts[2]
			}
		}

		output = append(output, addrLine)
	}

	return output, nil
}


func (s Service) SendLookups(lookups ...*street.Lookup) error {
	err := s.LookupClient.SendLookups(lookups...)
	if err != nil {
		return fmt.Errorf("error while sending lookups: %w", err)
	}

	return nil
}
