package model

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type Location struct {
	Lan float64 `json:"lan"`
	Lng float64 `json:"lng"`
}

func (c Location) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.Lan, v.Min(0)),
		v.Field(&c.Lng, v.Min(0)),
	)
}
