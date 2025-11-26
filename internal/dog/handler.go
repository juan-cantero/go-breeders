package dog

import (
	"net/http"

	"github.com/tsawler/toolbox"
)

// Handler handles HTTP requests for dog domain
type Handler struct {
	service *Service
}

// NewHandler creates a new dog handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetAllBreedsJSON returns all dog breeds as JSON
func (h *Handler) GetAllBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	breeds, err := h.service.GetAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, breeds)
}

// GetBreedByIDJSON returns a specific dog breed as JSON
func (h *Handler) GetBreedByIDJSON(w http.ResponseWriter, r *http.Request) {
	// Implementation would get ID from URL params
	// This is just a placeholder for now
}

// GetAllDogsJSON returns all dogs as JSON
func (h *Handler) GetAllDogsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	dogs, err := h.service.GetAllDogs()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogs)
}
