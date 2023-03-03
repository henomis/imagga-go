package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator"
)

const (
	faceGroupingsPath = "/faces/groupings"
)

type FacesGroupings struct {
	CallbackURL *string  `json:"callback_url,omitempty" validate:"omitempty,url"`
	Faces       []string `json:"faces,omitempty" validate:"-"`
}

func (f *FacesGroupings) Path() (string, error) {
	return faceGroupingsPath, nil
}

func (f *FacesGroupings) Validate() error {
	validate := validator.New()

	if err := validate.Struct(f); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return nil

}

func (f *FacesGroupings) Encode() (string, error) {

	if err := f.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	bytes, err := json.Marshal(f)
	if err != nil {
		return "", fmt.Errorf("failed to encode data: %w", err)
	}

	return string(bytes), nil

}
