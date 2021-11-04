package api

import (
	"github.com/2403905/location_history/internal/model"
)

type LocationPayload struct {
	OrderId string           `json:"order_id"`
	History []model.Location `json:"history"`
}
