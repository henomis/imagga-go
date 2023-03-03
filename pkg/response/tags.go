package response

import "io"

type Tags struct {
	Status `json:"status"`
	Result TagResult `json:"result"`
}

type TagResult struct {
	Tags []Tag `json:"tags,omitempty"`
}

type Tag struct {
	Confidence float64           `json:"confidence"`
	Origin     *string           `json:"origin,omitempty"`
	SynsetID   *string           `json:"synset_id,omitempty"`
	Tag        map[string]string `json:"tag"`
}

func (t *Tags) Decode(body io.ReadCloser) error {
	return decode(body, t)
}

func (t *Tags) SetBody(body io.ReadCloser) {}
