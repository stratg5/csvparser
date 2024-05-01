package addyverification

import street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"

type LookupSender interface {
	SendLookups(lookups ...*street.Lookup) error
}