package addylookup

import (
	"empora/entities"
	"fmt"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

// TODO rename this?
// TODO decide if we want to break out the low level stuff to another place/encapsulate it

type Client struct {
	LookupClient LookupSender
}

func NewClient(lookupSender LookupSender) Client {
	return Client{
		LookupClient: lookupSender,
	}
}

func (c Client) BuildLookups(addresses []entities.Address) []*street.Lookup {
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

func (c Client) SendLookups(lookups ...*street.Lookup) error {
	err := c.LookupClient.SendLookups(lookups...)
	if err != nil {
		return fmt.Errorf("error while sending lookups: %w", err)
	}

	return nil
}