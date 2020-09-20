package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jadugnap/golang-fpl-101/pkg/client"
	"github.com/jadugnap/golang-fpl-101/pkg/element"
	"github.com/jadugnap/golang-fpl-101/pkg/fpl"
)

func main() {
	start := time.Now()
	defer func() {
		log.Printf("Took %v overall to execute main()\n", time.Since(start))
	}()
	// define generic client
	genCli := client.GenericClient{
		HTTPClient: http.Client{Timeout: time.Second * 10},
	}

	// define global FPL instance
	fplInfo := fpl.FPL{
		Client: genCli,
	}
	fplInfo.Client.Endpoint = fpl.Endpoint
	// get bootstrap-static data
	fplInfo.GetFplResponseToCsv()
	if len(fplInfo.Res.Players) == 0 {
		log.Println("error executing getFplResponse().")
		return
	}

	// define global Element instance
	eInfo := element.Element{
		Client: genCli,
	}
	eInfo.Client.Endpoint = element.RawEndpoint
	// get necessary data from fplInfo
	for _, p := range fplInfo.Res.Players {
		eInfo.PlayerIDlist = append(eInfo.PlayerIDlist, p.ID)
	}
	// get element-summary data
	eInfo.GetElementSummaryToCsv()

	// get necessary data from eInfo
	fplInfo.Team2Gw2Points = eInfo.Team2Gw2Points
	fplInfo.ToCsv()
}
