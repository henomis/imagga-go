package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	categoriesPath = "/categories"
)

type Categories struct {
	ImageURL      string  `json:"image_url,omitempty" validate:"omitempty,url"`
	ImageUploadID string  `json:"image_upload_id,omitempty" validate:"-"`
	Language      *string `json:"language,omitempty" validate:"omitempty,oneof=ar bg bs en ca cs"`
	SaveIndex     *string `json:"save_index,omitempty" validate:"-"`
	SaveID        *string `json:"save_id,omitempty" validate:"-"`
	CategorizerID string  `json:"categorizer_id,omitempty" validate:"required"`
}

func (c *Categories) Path() (string, error) {

	categoriesParameters, err := c.Encode()

	if err != nil {
		return "", err
	}

	path := categoriesPath + "/" + c.CategorizerID

	return path + "?" + categoriesParameters, nil
}

func (c *Categories) Validate() error {
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

func (c *Categories) Encode() (string, error) {

	if err := c.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("image_url", &c.ImageURL)
	parameters.Add("image_upload_id", &c.ImageUploadID)
	parameters.Add("language", c.Language)
	parameters.Add("save_index", c.SaveIndex)
	parameters.Add("save_id", c.SaveID)

	return parameters.Encode(), nil
}
