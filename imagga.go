package imaggago

import (
	"fmt"
	"time"

	"github.com/henomis/imagga-go/internal/pkg/httpclient"
	"github.com/henomis/imagga-go/pkg/request"
	"github.com/henomis/imagga-go/pkg/response"
)

const (
	ImaggaEndpointV2 = "https://api.imagga.com/v2"
)

var (
	defaultHttpClientTimeout = 10 * time.Second
)

type ImaggaClient struct {
	httpClient *httpclient.HttpClient
}

func New(endpoint string, apiKey, apiSecret string, timeout *time.Duration) *ImaggaClient {

	if timeout == nil {
		timeout = &defaultHttpClientTimeout
	}

	return &ImaggaClient{
		httpClient: httpclient.New(endpoint, apiKey, apiSecret, *timeout),
	}
}

func (i *ImaggaClient) Tags(tagRequest *request.Tags) (*response.Tags, error) {

	tagResponse := &response.Tags{}
	err := i.httpClient.Get(tagRequest, tagResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return tagResponse, nil
}

func (i *ImaggaClient) Categorizers(categorizersRequest *request.Categorizers) (*response.Categorizers, error) {

	categorizersResponse := &response.Categorizers{}
	err := i.httpClient.Get(categorizersRequest, categorizersResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return categorizersResponse, nil
}

func (i *ImaggaClient) Categories(categorizersRequest *request.Categories) (*response.Categories, error) {

	categoriesResponse := &response.Categories{}
	err := i.httpClient.Get(categorizersRequest, categoriesResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return categoriesResponse, nil
}

func (i *ImaggaClient) Croppings(croppingsRequest *request.Croppings) (*response.Croppings, error) {

	croppingsResponse := &response.Croppings{}
	err := i.httpClient.Get(croppingsRequest, croppingsResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return croppingsResponse, nil
}

func (i *ImaggaClient) Colors(croppingsRequest *request.Colors) (*response.Colors, error) {

	colorsResponse := &response.Colors{}
	err := i.httpClient.Get(croppingsRequest, colorsResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return colorsResponse, nil
}

func (i *ImaggaClient) FacesDetections(facesDetectionsRequest *request.FacesDetections) (*response.FacesDetections, error) {

	facesDetectionsResponse := &response.FacesDetections{}
	err := i.httpClient.Get(facesDetectionsRequest, facesDetectionsResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return facesDetectionsResponse, nil
}

func (i *ImaggaClient) FacesSimilarity(facesSimilarityRequest *request.FacesSimilarity) (*response.FacesSimilarity, error) {

	facesSimilarityResponse := &response.FacesSimilarity{}
	err := i.httpClient.Get(facesSimilarityRequest, facesSimilarityResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return facesSimilarityResponse, nil
}

func (i *ImaggaClient) FacesGroupings(facesGroupingsRequest *request.FacesGroupings) (*response.FacesGroupings, error) {

	facesGroupingResponse := &response.FacesGroupings{}
	err := i.httpClient.Post(facesGroupingsRequest, facesGroupingResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return facesGroupingResponse, nil
}

func (i *ImaggaClient) Tickets(ticketsRequest *request.Tickets) (*response.Tickets, error) {

	ticketsResponse := &response.Tickets{}
	err := i.httpClient.Get(ticketsRequest, ticketsResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return ticketsResponse, nil
}

func (i *ImaggaClient) Text(textRequest *request.Text) (*response.Text, error) {

	textResponse := &response.Text{}
	err := i.httpClient.Get(textRequest, textResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return textResponse, nil
}

func (i *ImaggaClient) Uploads(uploadsRequest *request.Uploads) (*response.Uploads, error) {

	uploadsResponse := &response.Uploads{}
	err := i.httpClient.Upload(uploadsRequest, uploadsResponse, uploadsRequest.FilePath)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return uploadsResponse, nil
}

func (i *ImaggaClient) DeleteUploads(uploadsRequest *request.Uploads) (*response.Uploads, error) {

	uploadsResponse := &response.Uploads{}
	err := i.httpClient.Delete(uploadsRequest, uploadsResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return uploadsResponse, nil
}

// TODO: Batches

func (i *ImaggaClient) Usage(usageRequest *request.Usage) (*response.Usage, error) {

	usageResponse := &response.Usage{}
	err := i.httpClient.Get(usageRequest, usageResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return usageResponse, nil
}

func (i *ImaggaClient) Barcodes(barcodesRequest *request.Barcodes) (*response.Barcodes, error) {

	barcodesResponse := &response.Barcodes{}
	err := i.httpClient.Get(barcodesRequest, barcodesResponse)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return barcodesResponse, nil
}
