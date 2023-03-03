package response

import "io"

type Categorizers struct {
	Status `json:"status"`
	Result CategorizersResult `json:"result"`
}

type CategorizersResult struct {
	Categorizers []Categorizer `json:"categorizers"`
}

type Categorizer struct {
	ID     string   `json:"id"`
	Labels []string `json:"labels"`
	Title  string   `json:"title"`
}

func (c *Categorizers) Decode(body io.ReadCloser) error {
	return decode(body, c)
}

func (c *Categorizers) SetBody(body io.ReadCloser) {}
