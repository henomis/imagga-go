package response

import "io"

type Text struct {
	Status `json:"status"`
	Result TextResult `json:"result"`
}

type TextResult struct {
	Text []struct {
		Data        string      `json:"data"`
		Coordinates Coordinates `json:"coordinates"`
	} `json:"text"`
}

func (t *Text) Decode(body io.ReadCloser) error {
	return decode(body, t)
}

func (t *Text) SetBody(body io.ReadCloser) {}
