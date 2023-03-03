package response

import "io"

type Croppings struct {
	Status    `json:"status"`
	Result    CroppingsResult `json:"result"`
	imageData io.ReadCloser
}

type CroppingsResult struct {
	Croppings []Cropping `json:"croppings"`
}

type Cropping struct {
	TargetHeight int `json:"target_height"`
	TargetWidth  int `json:"target_width"`
	X1           int `json:"x1"`
	X2           int `json:"x2"`
	Y1           int `json:"y1"`
	Y2           int `json:"y2"`
}

func (c *Croppings) Decode(body io.ReadCloser) error {
	return decode(body, c)
}

func (c *Croppings) SetBody(body io.ReadCloser) {
	c.imageData = body
}

func (c *Croppings) Image() io.ReadCloser {
	return c.imageData
}

func (c *Croppings) IsSuccess() bool {

	if c.imageData != nil {
		return true
	}

	return c.Status.IsSuccess()
}
