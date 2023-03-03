package main

import (
	"io"
	"log"
	"os"

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

	response, err := imaggaClient.Croppings(
		&request.Croppings{
			ImageURL:    "https://imagga.com/static/images/tagging/wind-farm-538576_640.jpg",
			Resolution:  "100x100",
			ImageResult: newBool(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if !response.IsSuccess() {
		log.Fatal(response.Error())
	}

	image := response.Image()
	defer image.Close()

	file, err := os.Create("croppings.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := io.Copy(file, image); err != nil {
		log.Fatal(err)
	}

}

// support methods

func newBool(value bool) *bool {
	return &value
}
