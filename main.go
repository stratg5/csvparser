package main

import (
	"csvparser/address"
	"csvparser/csv"
	"csvparser/driver"
	"csvparser/entities"
	"csvparser/ostools"
	"flag"

	"github.com/smartystreets/smartystreets-go-sdk/wireup"
)

// the main file is just for initialization. everything gets injected into the driver service
// the driver service coordinates the individual pieces
// this makes mocking and testing much easier and more complete
func main() {
	var inputPath = flag.String("inputPath", "./input.csv", "the path to the csv, defaults to ./input.csv")
	var outputPath = flag.String("outputPath", "./output.csv", "the path to the csv, defaults to ./output.csv")
	var idFlag = flag.String("apiID", "", "the api ID")
	var tokenFlag = flag.String("apiToken", "", "the api token")

	flag.Parse()

	id := entities.ID
	token := entities.Token

	if *idFlag != "" {
		id = *idFlag
	}

	if *tokenFlag != "" {
		token = *tokenFlag
	}

	addressSvc := address.NewService(wireup.BuildUSStreetAPIClient(wireup.SecretKeyCredential(id, token)))
	osToolSvc := ostools.NewService()
	csvSvc := csv.NewService(osToolSvc)

	driverSvc := driver.NewService(addressSvc, csvSvc)
	err := driverSvc.ParseCSVAndGenerateOutput(*inputPath, *outputPath)
	if err != nil {
		println(err.Error())
	}
}
