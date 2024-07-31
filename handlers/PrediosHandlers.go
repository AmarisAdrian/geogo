package handlers

import (
	"api/models"
	"api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PredioHandler struct {
	Repo repository.PredioRepository
}

func (h *PredioHandler) GetPredios(w http.ResponseWriter, r *http.Request) {
	predio, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(predio)
}
func (h *PredioHandler) GetPredio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	predioId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID Predio invalido", http.StatusBadRequest)
		return
	}
	predio, err := h.Repo.GetById(predioId)
	if err != nil {
		http.Error(w, "Predio no encontrado", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(predio)
}
func (h *PredioHandler) CreatePredio(w http.ResponseWriter, r *http.Request) {
	var newPredio models.PrediosModel
	if err := json.NewDecoder(r.Body).Decode(&newPredio).Error; err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdPredio, err := h.Repo.Create(newPredio)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPredio)
}
func (h *PredioHandler) UpdatePredio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	predioId, err := strconv.Atoi(vars["t_id"])
	if err != nil {
		http.Error(w, "ID Predio invalido", http.StatusBadRequest)
		return
	}
	var updatePredio models.PrediosModel
	if err := json.NewDecoder(r.Body).Decode(&updatePredio); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}
	predio, err := h.Repo.Update(predioId, updatePredio)
	if err != nil {
		http.Error(w, "Predio no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(predio)
}

func (h *PredioHandler) DeletePredio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	predioId, err := strconv.Atoi(vars["t_id"])
	if err != nil {
		http.Error(w, "ID Predio invalido", http.StatusBadRequest)
		return
	}
	if err := h.Repo.Delete(predioId); err != nil {
		http.Error(w, "Predio no encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
