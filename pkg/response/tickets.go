package response

import "io"

type Tickets struct {
	Status `json:"status"`
	Result TicketsResult `json:"result"`
}

type TicketsResult struct {
	IsFinal      bool `json:"is_final"`
	TicketResult struct {
		Groups []string `json:"groups,omitempty"`
	} `json:"ticket_result,omitempty"`
}

func (t *Tickets) Decode(body io.ReadCloser) error {
	return decode(body, t)
}

func (t *Tickets) SetBody(body io.ReadCloser) {}
