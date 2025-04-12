package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
	"go-rest-api/pkg/logger"
)

type ItemHandler struct {
	Service *services.ItemService
}

func NewItemHandler(itemService *services.ItemService) *ItemHandler {
	return &ItemHandler{Service: itemService}
}

func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := h.Service.GetAllItems()
	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *ItemHandler) GetItemById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "invalid ID", "")
		return
	}
	item, found := h.Service.GetItemById(id)
	if !found {
		utils.RespondWithError(w, http.StatusFound, "item not found", "")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, item)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Name == "" {
		logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "invalid input", "")
		return
	}
	item := h.Service.CraeteItem(input.Name)
	utils.RespondWithJSON(w, http.StatusCreated, item)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "invalid ID", "")
		return
	}
	if !h.Service.DeleteItem(id) {
		utils.RespondWithError(w, http.StatusFound, "item not found", "")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
