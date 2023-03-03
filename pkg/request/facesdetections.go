package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/imagga-go/internal/pkg/parameters"
)

const (
	faceDetectionsPath = "/faces/detections"
)

type FacesDetections struct {
	ImageURL      string `json:"image_url,omitempty" validate:"omitempty,url"`
	ImageUploadID string `json:"image_upload_id,omitempty" validate:"-"`
	ReturnFaceID  *bool  `json:"return_face_id,omitempty" validate:"-"`
}

func (f *FacesDetections) Path() (string, error) {

	facesDetectionsParameters, err := f.Encode()

	if err != nil {
		return "", err
	}

	return faceDetectionsPath + "?" + facesDetectionsParameters, nil
}

func (f *FacesDetections) Validate() error {
	validate := validator.New()

	if err := validate.Struct(f); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"image_url":       f.ImageURL,
			"image_upload_id": f.ImageUploadID,
		},
	)

}

func (f *FacesDetections) Encode() (string, error) {

	if err := f.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	parameters := parameters.New()
	parameters.Add("image_url", &f.ImageURL)
	parameters.Add("image_upload_id", &f.ImageUploadID)
	parameters.AddBoolAsInt("return_face_id", f.ReturnFaceID)

	return parameters.Encode(), nil
}
