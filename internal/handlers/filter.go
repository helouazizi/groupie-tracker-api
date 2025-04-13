package handlers

import (
	"encoding/json"
	"net/http"

	"go-rest-api/internal/models"
	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
	"go-rest-api/pkg/logger"
)

type FilterHandler struct {
	Service *services.FilterService
}

func NewFilterHandler(filterservice *services.FilterService) *FilterHandler {
	return &FilterHandler{Service: filterservice}
}

func (h *FilterHandler) Filter(w http.ResponseWriter, r *http.Request) {
	var data models.FilterRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "bad request", "some data it incorrect or empty")
		return
	}
	artists, err := h.Service.Filter(data)
	if err != nil {
		logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "bad request", "some data it incorrect or empty")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, artists)
}
