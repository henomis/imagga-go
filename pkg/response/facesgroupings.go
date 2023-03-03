package response

import "io"

type FacesGroupings struct {
	Status `json:"status"`
	Result FacesGroupingsResult `json:"result"`
}

type FacesGroupingsResult struct {
	TicketID string `json:"ticket_id"`
}

func (f *FacesGroupings) Decode(body io.ReadCloser) error {
	return decode(body, f)
}

func (f *FacesGroupings) SetBody(body io.ReadCloser) {}
