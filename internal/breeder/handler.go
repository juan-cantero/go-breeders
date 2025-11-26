package breeder

import (
	"net/http"

	"github.com/tsawler/toolbox"
)

// Handler handles HTTP requests for breeder domain
type Handler struct {
	service *Service
}

// NewHandler creates a new breeder handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetAllBreedersJSON returns all breeders as JSON
func (h *Handler) GetAllBreedersJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	breeders, err := h.service.GetAllBreeders()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, breeders)
}
