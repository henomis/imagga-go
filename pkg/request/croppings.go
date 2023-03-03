package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	croppingsPath = "/croppings"
)

type Croppings struct {
	ImageURL       string   `json:"image_url,omitempty" validate:"omitempty,url"`
	ImageUploadID  string   `json:"image_upload_id,omitempty" validate:"-"`
	Resolution     string   `json:"resolution" validate:"required"`
	NoScaling      *bool    `json:"no_scaling,omitempty" validate:"-"`
	RectPercentage *float64 `json:"rect_percentage,omitempty" validate:"omitempty,gte=0,lte=1"`
	ImageResult    *bool    `json:"image_result,omitempty" validate:"-"`
}

func (c *Croppings) Path() (string, error) {

	croppingsParameters, err := c.Encode()

	if err != nil {
		return "", err
	}

	return croppingsPath + "?" + croppingsParameters, nil
}

func (c *Croppings) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"image_url":       c.ImageURL,
			"image_upload_id": c.ImageUploadID,
		},
	)

}

func (c *Croppings) Encode() (string, error) {

	if err := c.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("image_url", &c.ImageURL)
	parameters.Add("image_upload_id", &c.ImageUploadID)
	parameters.Add("resolution", &c.Resolution)
	parameters.AddBoolAsInt("no_scaling", c.NoScaling)
	parameters.AddFloat("rect_percentage", c.RectPercentage)
	parameters.AddBoolAsInt("image_result", c.ImageResult)

	return parameters.Encode(), nil
}
