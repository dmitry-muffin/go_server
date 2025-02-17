package handlers

import (
	"encoding/json"
	"go_server/internal/storage"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	Store *storage.ReadyStorage
}

func NewHandler(store *storage.ReadyStorage) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) AddUserHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user storage.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info("request body decoded", slog.Any("request", user))

		if err := h.Store.AddUser(user); err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			log.Error("failed to add User to storage", slog.Any("user:", user))
		}

		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Error("failed to response", err.Error())
			return
		}
	}
}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	user, err := h.Store.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
