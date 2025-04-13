package handlers

import (
	"net/http"

	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
)

type AllArtistsHandler struct {
	Service *services.AllArtistsService
}

func NewAllArtistsHandler(allartistservice *services.AllArtistsService) *AllArtistsHandler {
	return &AllArtistsHandler{Service: allartistservice}
}

func (a *AllArtistsHandler) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	artists := a.Service.GetAllArtists()
	utils.RespondWithJSON(w, http.StatusOK, artists)
}
