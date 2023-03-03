package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	usagePath = "/usage"
)

type Usage struct {
	History     *bool `json:"history,omitempty" validate:"-"`
	Concurrency *bool `json:"concurrency,omitempty" validate:"-"`
}

func (u *Usage) Path() (string, error) {

	usageParameters, err := u.Encode()

	if err != nil {
		return "", err
	}

	return usagePath + "?" + usageParameters, nil
}

func (u *Usage) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return nil
}

func (u *Usage) Encode() (string, error) {

	if err := u.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.AddBoolAsInt("history", u.History)
	parameters.AddBoolAsInt("concurrency", u.Concurrency)

	return parameters.Encode(), nil
}
