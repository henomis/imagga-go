package response

import "io"

type Uploads struct {
	Status `json:"status"`
	Result UploadsResult `json:"result"`
}

type UploadsResult struct {
	UploadID string `json:"upload_id"`
}

func (u *Uploads) Decode(body io.ReadCloser) error {
	return decode(body, u)
}

func (u *Uploads) SetBody(body io.ReadCloser) {}
