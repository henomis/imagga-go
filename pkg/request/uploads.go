package request

import (
	"fmt"

	"github.com/go-playground/validator"
)

const (
	uploadsPath = "/uploads"
)

type Uploads struct {
	UploadID string `json:"upload_id,omitempty" validate:"-"`
	FilePath string `json:"file_path,omitempty" validate:"-"`
}

func (u *Uploads) Path() (string, error) {

	path := uploadsPath
	if len(u.UploadID) > 0 {
		path += "/" + u.UploadID
	}

	return path, nil
}

func (u *Uploads) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"upload_id": u.UploadID,
			"file_path": u.FilePath,
		},
	)
}

func (u *Uploads) Encode() (string, error) {
	return "", nil
}
