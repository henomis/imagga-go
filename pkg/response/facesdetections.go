package response

import "io"

type FacesDetections struct {
	Status `json:"status"`
	Result FacesDetectionsResult `json:"result"`
}

type FacesDetectionsResult struct {
	Faces []Face `json:"faces"`
}

type Face struct {
	Confidence  float64     `json:"confidence"`
	Coordinates Coordinates `json:"coordinates"`
	FaceID      string      `json:"face_id"`
}

type Coordinates struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Xmax   int `json:"xmax"`
	Xmin   int `json:"xmin"`
	Ymax   int `json:"ymax"`
	Ymin   int `json:"ymin"`
}

func (f *FacesDetections) Decode(body io.ReadCloser) error {
	return decode(body, f)
}

func (f *FacesDetections) SetBody(body io.ReadCloser) {}
