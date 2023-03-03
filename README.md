# Imagga SDK for Go


[![Build Status](https://github.com/henomis/imagga-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/henomis/imagga-go/actions/workflows/test.yml?query=branch%3Amain) [![GoDoc](https://godoc.org/github.com/henomis/imagga-go?status.svg)](https://godoc.org/github.com/henomis/imagga-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/imagga-go)](https://goreportcard.com/report/github.com/henomis/imagga-go) [![GitHub release](https://img.shields.io/github/release/henomis/imagga-go.svg)](https://github.com/henomis/imagga-go/releases)

This is Imagga's **unofficial** Go client, designed to enable you to use Imagga's services easily from your own applications.

## Imagga

Imagga is a cloud-based analytics service that through APIs allows you extract meaning from images and text.

## SDK versions

API version | SDK version
--- | ---
v2 | v1.0.0


## Getting started

### Installation

You can load imagga-go into your project by using:
```
go get github.com/henomis/imagga-go
```


### Configuration

The only thing you need to start using Imagga's APIs is the developer license key and related secret. Copy it and paste it in the corresponding place in the code, select the API you want to use and the parameters you want to use, and that's it.


### Usage

Please refer to the [examples folder](examples/cmd/) to see how to use the SDK.

Here below a simple usage example:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	imaggago "github.com/henomis/imagga-go"
	"github.com/henomis/imagga-go/pkg/request"
)

const apiKey = "YOUR_API_KEY"
const apiSecret = "YOUR_API_SECRET"

func main() {

	imaggaClient := imaggago.New(
		imaggago.ImaggaEndpointV2,
		apiKey,
		apiSecret,
		nil,
	)

	response, err := imaggaClient.Barcodes(
		&request.Barcodes{
			ImageURL: "https://imagga.com/static/images/technology/barcode.png",
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
```