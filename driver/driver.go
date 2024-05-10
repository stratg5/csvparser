package driver

import (
	"csvparser/address"
	"csvparser/csv"
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

func (s Service) ParseCSVAndGenerateOutput(inputPath, outputPath string) error {
	rawData, err := s.csvSvc.ReadCSV(inputPath)
	if err != nil {
		return fmt.Errorf("error while parsing csv: %w", err)
	}

	addresses := s.addressSvc.BuildAddressesFromRawData(rawData)

	lookups := s.addressSvc.BuildLookupsFromAddresses(addresses)

	err = s.addressSvc.SendLookups(lookups...)
	if err != nil {
		return fmt.Errorf("error while sending lookups: %w", err)
	}

	outputData, err := s.addressSvc.BuildRawDataFromLookups(addresses, lookups)
	if err != nil {
		return fmt.Errorf("error while building raw data: %w", err)
	}

	err = s.csvSvc.WriteCSV(outputPath, outputData)
	if err != nil {
		return fmt.Errorf("error while writing output to csv: %w", err)
	}

	return nil
}
