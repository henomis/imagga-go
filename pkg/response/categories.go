package response

import "io"

type Categories struct {
	Status `json:"status"`
	Result CategoriesResult `json:"result"`
}

type CategoriesResult struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Confidence float64 `json:"confidence"`
	Name       Name    `json:"name"`
}

type Name struct {
	En string `json:"en"`
}

func (c *Categories) Decode(body io.ReadCloser) error {
	return decode(body, c)
}

func (c *Categories) SetBody(body io.ReadCloser) {}
