package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type Status struct {
	StatusCode int    `json:"status_code"`
	Text       string `json:"text"`
	Type       string `json:"type"`
}

func (s *Status) IsSuccess() bool {
	return (s.StatusCode >= 200 && s.StatusCode < 300) || s.Type == "success"
}

func (s *Status) Error() error {
	return fmt.Errorf("%s", s.Text)
}

func (s *Status) SetStatusCode(code int) {
	s.StatusCode = code
}

func decode(body io.ReadCloser, data interface{}) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(data)
}
