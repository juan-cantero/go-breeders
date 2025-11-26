package cat

import (
	"net/http"

	"github.com/tsawler/toolbox"
)

// Handler handles HTTP requests for cat domain
type Handler struct {
	service *Service
}

// NewHandler creates a new cat handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetAllBreedsJSON returns all cat breeds as JSON
func (h *Handler) GetAllBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	breeds, err := h.service.GetAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, breeds)
}

// GetAllCatsJSON returns all cats as JSON
func (h *Handler) GetAllCatsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	cats, err := h.service.GetAllCats()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cats)
}
