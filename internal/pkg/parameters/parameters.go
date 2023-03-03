package parameters

import (
	"fmt"
	"net/url"
)

type Parameters url.Values

func New() *Parameters {
	return &Parameters{}
}

func (p *Parameters) Add(key string, value *string) {
	if value != nil && *value != "" {
		(*url.Values)(p).Add(key, *value)
	}
}

func (p *Parameters) Encode() string {
	return (*url.Values)(p).Encode()
}

func (p *Parameters) AddInt(key string, value *int) {
	if value != nil {
		(*url.Values)(p).Add(key, fmt.Sprintf("%d", *value))
	}
}

func (p *Parameters) AddBool(key string, value *bool) {
	if value != nil {
		(*url.Values)(p).Add(key, fmt.Sprintf("%t", *value))
	}
}

func (p *Parameters) AddBoolAsInt(key string, value *bool) {
	if value != nil {
		if *value {
			(*url.Values)(p).Add(key, "1")
		} else {
			(*url.Values)(p).Add(key, "0")
		}
	}
}

func (p *Parameters) AddFloat(key string, value *float64) {
	if value != nil {
		(*url.Values)(p).Add(key, fmt.Sprintf("%f", *value))
	}
}
