package address

import street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"

type MockLookupSender struct {
	SendLookupsFn func(lookups ...*street.Lookup) error
}

func (m MockLookupSender) SendLookups(lookups ...*street.Lookup) error {
	return m.SendLookupsFn(lookups...)
}