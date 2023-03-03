package response

import "io"

type Barcodes struct {
	Status `json:"status"`
	Result BarcodesResult `json:"result"`
}

type BarcodesResult struct {
	Barcodes []Barcode `json:"barcodes"`
}

type Barcode struct {
	Data string `json:"data"`
	Type string `json:"type"`
	X1   int    `json:"x1"`
	X2   int    `json:"x2"`
	Y1   int    `json:"y1"`
	Y2   int    `json:"y2"`
}

func (b *Barcodes) Decode(body io.ReadCloser) error {
	return decode(body, b)
}

func (b *Barcodes) SetBody(body io.ReadCloser) {}
