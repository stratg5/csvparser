package driver

import (
	"empora/address"
	"empora/csv"
)

type Service struct {
	addressSvc address.AddressProvider
	csvSvc     csv.CSVProvider
}

func NewService(addressSvc address.AddressProvider, csvSvc csv.CSVProvider) Service {
	return Service{
		addressSvc: addressSvc,
		csvSvc:     csvSvc,
	}
}

func (s Service) ParseCSVAndGenerateOutput(path string) {
	// TODO
	// Send the fields to the endpoint for verification
	// Generate the output as a csv or as console output (flag for options)
}
