package main

import (
	"encoding/json"
	"fmt"
	"log"

	imaggago "github.com/henomis/imagga-go"
	"github.com/henomis/imagga-go/pkg/request"
)

const apiKey = "YOUR-API-KEY"
const apiSecret = "YOUR-API-SECRET"

func main() {

	imaggaClient := imaggago.New(
		imaggago.ImaggaEndpointV2,
		apiKey,
		apiSecret,
		nil,
	)

	response, err := imaggaClient.Tickets(
		&request.Tickets{
			TicketID: "82eb33f78744ae3f7a26fde18389c8adaf470d03896489e84eeefe48f2e2a187b1dc9d4b55908b391267d271a510d92ed98d",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if !response.IsSuccess() {
		log.Fatal(response.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}
