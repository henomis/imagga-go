package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	textPath = "/text"
)

type Text struct {
	ImageURL string `json:"image_url" validate:"url"`
}

func (t *Text) Path() (string, error) {

	textParameters, err := t.Encode()

	if err != nil {
		return "", err
	}

	return textPath + "?" + textParameters, nil
}

func (t *Text) Validate() error {
	validate := validator.New()

	if err := validate.Struct(t); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return nil

}

func (t *Text) Encode() (string, error) {

	if err := t.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("image_url", &t.ImageURL)

	return parameters.Encode(), nil
}
