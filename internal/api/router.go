// Package api Location History API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//
// swagger:meta
package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/2403905/location_history/internal/service"
)

func NewHandler(locationService *service.Location) http.Handler {
	reportHandler := NewReportHandler(locationService)
	router := mux.NewRouter()
	router.HandleFunc("/location/{order_id}/now", reportHandler.AddLocation).Methods(http.MethodPost).
		HeadersRegexp("Content-Type", "application/json")

	router.HandleFunc("/location/{order_id}", reportHandler.GetLocation).Methods(http.MethodGet)

	router.HandleFunc("/location/{order_id}", reportHandler.DeleteLocation).Methods(http.MethodDelete)

	return router
}
