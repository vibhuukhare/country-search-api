package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vibhu.khare/country-api/services"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(service *services.CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (h *CountryHandler) SearchCountryName(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(rw, "Missing name query parameter", http.StatusBadRequest)
		return
	}

	country, err := h.service.GetCountryData(name)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(country)

}
