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

	response, err := imaggaClient.FacesSimilarity(
		&request.FacesSimilarity{
			FaceID:       "1ede163690e7a6b3a2033c694bfc1319ff9cb24f491a44fdfc7d45ff2c74e9bd",
			SecondFaceID: "645649873bbbe1ae31dd2c8c8714000e41d4d76e5ac24b5acb95c6ce16f09fbe",
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
