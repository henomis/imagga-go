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

	response, err := imaggaClient.Tags(
		&request.Tags{
			ImageURL:  "https://imagga.com/static/images/tagging/wind-farm-538576_640.jpg",
			Verbose:   newBool(true),
			Limit:     newInt(10),
			Threshold: newFloat64(0.5),
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

// Support methods

func newString(s string) *string {
	return &s
}

func newInt(i int) *int {
	return &i
}

func newFloat64(f float64) *float64 {
	return &f
}

func newBool(b bool) *bool {
	return &b
}
