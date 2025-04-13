package handlers

import (
	"net/http"
	"strconv"

	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
	"go-rest-api/pkg/logger"
)

type ArtistsDetailsHandler struct {
	Service *services.ArtistsDetailsService
}

func NewArtistDetailsService(service *services.ArtistsDetailsService) *ArtistsDetailsHandler {
	return &ArtistsDetailsHandler{Service: service}
}

func (s *ArtistsDetailsHandler) GetArtistDetails(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "invalid ID", "invalid artist id")
		return
	}
	details, err := s.Service.GetArtistDetails(id)
	if err != nil {
		logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "artist not found", "artist my be deleted or some thing else")
		return
	}
	utils.RespondWithJSON(w, http.StatusFound, details)
}
