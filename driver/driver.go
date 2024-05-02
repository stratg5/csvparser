package driver

import (
	"empora/address"
	"empora/csv"
	"fmt"
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

func (s Service) ParseCSVAndGenerateOutput(path string) error {
	rawData, err := s.csvSvc.ReadCSV(path)
	if err != nil {
		return fmt.Errorf("error while parsing csv: %w", err)
	}

	addresses := s.addressSvc.BuildAddressesFromRawData(rawData)

	lookups := s.addressSvc.BuildLookups(addresses)

	err = s.addressSvc.SendLookups(lookups...)
	if err != nil {
		return fmt.Errorf("error while sending lookups: %w", err)
	}
	
	// TODO
	// Generate the output as a csv or as console output (flag for options)
	return nil
}
