package handler

import (
	"encoding/json"
	"net/http"
	"one-lab/internal/model"

	"github.com/google/uuid"
)

func (h *Handler) addUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var users model.User
	err := decoder.Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	users.ID = id
	h.services.Add(id, users.UserEntity)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users.ID)
}
