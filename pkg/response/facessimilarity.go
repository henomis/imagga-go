package response

import "io"

type FacesSimilarity struct {
	Status `json:"status"`
	Result FacesSimilarityResult `json:"result"`
}

type FacesSimilarityResult struct {
	Score float64 `json:"score"`
}

func (f *FacesSimilarity) Decode(body io.ReadCloser) error {
	return decode(body, f)
}

func (f *FacesSimilarity) SetBody(body io.ReadCloser) {}
