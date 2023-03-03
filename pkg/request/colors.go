package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	colorsPath = "/colors"
)

type Colors struct {
	ImageURL             string  `json:"image_url,omitempty" validate:"omitempty,url"`
	ImageUploadID        string  `json:"image_upload_id,omitempty" validate:"-"`
	ExtractOverallColors *bool   `json:"extract_overall_colors,omitempty" validate:"-"`
	ExtractObjectColors  *bool   `json:"extract_object_colors,omitempty" validate:"-"`
	OverallCount         *int    `json:"overall_count,omitempty" validate:"-"`
	SeparatedCount       *int    `json:"separated_count,omitempty" validate:"-"`
	Deterministic        *bool   `json:"deterministic,omitempty" validate:"-"`
	SaveIndex            *string `json:"save_index,omitempty" validate:"-"`
	SaveID               *string `json:"save_id,omitempty" validate:"-"`
	FeaturesType         *string `json:"features_type,omitempty" validate:"omitempty,oneof=overall object"`
}

func (c *Colors) Path() (string, error) {

	tagsParameters, err := c.Encode()

	if err != nil {
		return "", err
	}

	return colorsPath + "?" + tagsParameters, nil
}

func (c *Colors) Validate() error {
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

func (c *Colors) Encode() (string, error) {

	if err := c.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("image_url", &c.ImageURL)
	parameters.Add("image_upload_id", &c.ImageUploadID)
	parameters.AddBoolAsInt("extract_overall_colors", c.ExtractOverallColors)
	parameters.AddBoolAsInt("extract_object_colors", c.ExtractObjectColors)
	parameters.AddInt("overall_count", c.OverallCount)
	parameters.AddInt("separated_count", c.SeparatedCount)
	parameters.AddBoolAsInt("deterministic", c.Deterministic)
	parameters.Add("save_index", c.SaveIndex)
	parameters.Add("save_id", c.SaveID)
	parameters.Add("features_type", c.FeaturesType)

	return parameters.Encode(), nil
}
