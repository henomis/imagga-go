package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	tagsPath = "/tags"
)

type Tags struct {
	ImageURL        string   `json:"image_url,omitempty" validate:"omitempty,url"`
	ImageUploadID   string   `json:"image_upload_id,omitempty" validate:"-"`
	Language        *string  `json:"language,omitempty" validate:"omitempty,oneof=ar bg bs en ca cs"`
	Verbose         *bool    `json:"verbose,omitempty"`
	Limit           *int     `json:"limit,omitempty"`
	Threshold       *float64 `json:"threshold,omitempty"`
	DecreaseParents *bool    `json:"decrease_parents,omitempty"`
	TaggerID        *string  `json:"tagger_id,omitempty"`
}

func (t *Tags) Path() (string, error) {

	tagsParameters, err := t.Encode()

	if err != nil {
		return "", err
	}

	path := tagsPath
	if t.TaggerID != nil {
		path += "/" + *t.TaggerID
	}

	return path + "?" + tagsParameters, nil
}

func (t *Tags) Validate() error {
	validate := validator.New()

	if err := validate.Struct(t); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"image_url":       t.ImageURL,
			"image_upload_id": t.ImageUploadID,
		},
	)

}

func (t *Tags) Encode() (string, error) {

	if err := t.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("image_url", &t.ImageURL)
	parameters.Add("image_upload_id", &t.ImageUploadID)
	parameters.Add("language", t.Language)
	parameters.AddBoolAsInt("verbose", t.Verbose)
	parameters.AddInt("limit", t.Limit)
	parameters.AddFloat("threshold", t.Threshold)
	parameters.AddBoolAsInt("decrease_parents", t.DecreaseParents)

	return parameters.Encode(), nil
}
