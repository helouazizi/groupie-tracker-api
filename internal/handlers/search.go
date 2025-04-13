package handlers

import (
	"net/http"

	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
)

type SearchHandler struct {
	Service *services.SearchService
}

func NewSearchHandler(filterservice *services.SearchService) *SearchHandler {
	return &SearchHandler{Service: filterservice}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("find")
	data := h.Service.Search(input)
	utils.RespondWithJSON(w, http.StatusOK, data)
}
