package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	barcodesPath = "/barcodes"
)

type Barcodes struct {
	ImageURL      string `json:"image_url,omitempty" validate:"omitempty,url"`
	ImageUploadID string `json:"image_upload_id,omitempty" validate:"-"`
}

func (b *Barcodes) Path() (string, error) {

	barcodesParameters, err := b.Encode()

	if err != nil {
		return "", err
	}

	return barcodesPath + "?" + barcodesParameters, nil
}

func (b *Barcodes) Validate() error {
	validate := validator.New()

	if err := validate.Struct(b); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"image_url":       b.ImageURL,
			"image_upload_id": b.ImageUploadID,
		},
	)

}

func (b *Barcodes) Encode() (string, error) {

	if err := b.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("image_url", &b.ImageURL)
	parameters.Add("image_upload_id", &b.ImageUploadID)

	return parameters.Encode(), nil
}
