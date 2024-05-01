package addyverification

import (
	"empora/entities"
	"fmt"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

// TODO rename this?
// TODO decide if we want to break out the low level stuff to another place/encapsulate it

type Service struct {
	LookupClient LookupSender
}

func NewService(lookupSender LookupSender) Service {
	return Service{
		LookupClient: lookupSender,
	}
}

func (s Service) BuildLookups(addresses []entities.Address) []*street.Lookup {
	lookups := []*street.Lookup{}
	for _, address := range addresses {
		lookups = append(lookups, &street.Lookup{
			Street: address.Street,
			City: address.Street,
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