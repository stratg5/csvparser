package main

import (
	"empora/address"
	"empora/csv"
	"empora/driver"
	"empora/entities"
	"empora/ostools"
	"flag"

	"github.com/smartystreets/smartystreets-go-sdk/wireup"
)

// the main file is just for initialization. everything gets injected into the driver service
// the driver service coordinates the individual pieces
// this makes mocking and testing much easier and more complete
func main() {
	var csvPath = flag.String("csvPath", "./addresses.csv", "the path to the csv, defaults to ./addresses.csv")
	var id = flag.String("apiID", "", "the api ID")
	var token = flag.String("apiToken", "", "the api token")

	println(*id)
	println(*token)

	// TODO check for the id and token in entities file, if not present then prompt

	addressSvc := address.NewService(wireup.BuildUSStreetAPIClient(wireup.SecretKeyCredential(entities.ID, entities.Token)))
	osToolSvc := ostools.NewService()
	csvSvc := csv.NewService(osToolSvc)

	driverSvc := driver.NewService(addressSvc, csvSvc)
	driverSvc.ParseCSVAndGenerateOutput(*csvPath)
}
