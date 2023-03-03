package request

import (
	"fmt"

	"github.com/go-playground/validator"
)

const (
	ticketsPath = "/tickets"
)

type Tickets struct {
	TicketID string `json:"ticket_id" validate:"required"`
}

func (f *Tickets) Path() (string, error) {
	return ticketsPath + "/" + f.TicketID, nil
}

func (f *Tickets) Validate() error {
	validate := validator.New()

	if err := validate.Struct(f); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return nil
}

func (f *Tickets) Encode() (string, error) {

	if err := f.Validate(); err != nil {
		return "", fmt.Errorf("invalid data: %w", err)
	}

	return "", nil
}
