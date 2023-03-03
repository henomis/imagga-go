package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	faceSimilarityPath = "/faces/similarity"
)

type FacesSimilarity struct {
	FaceID       string `json:"face_id" validate:"required"`
	SecondFaceID string `json:"second_face_id" validate:"required"`
}

func (f *FacesSimilarity) Path() (string, error) {

	facesSimilarityParameters, err := f.Encode()

	if err != nil {
		return "", err
	}

	return faceSimilarityPath + "?" + facesSimilarityParameters, nil
}

func (f *FacesSimilarity) Validate() error {
	validate := validator.New()

	if err := validate.Struct(f); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return nil

}

func (f *FacesSimilarity) Encode() (string, error) {

	if err := f.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("face_id", &f.FaceID)
	parameters.Add("second_face_id", &f.SecondFaceID)

	return parameters.Encode(), nil
}
