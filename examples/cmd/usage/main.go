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

	response, err := imaggaClient.Usage(
		&request.Usage{
			History:     newBool(true),
			Concurrency: newBool(true),
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

// support methods

func newBool(value bool) *bool {
	return &value
}