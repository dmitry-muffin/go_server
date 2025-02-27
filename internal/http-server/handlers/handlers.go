package handlers

import (
	"encoding/json"
	"go_server/internal/domain"
	"go_server/internal/storage"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	store *storage.ReadyStorage
	log   *slog.Logger
}

func NewHandler(store *storage.ReadyStorage, log *slog.Logger) *Handler {
	return &Handler{
		store: store,
		log:   log,
	}
}

func (h *Handler) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.log.Info("request body decoded", slog.Any("request", user))

	if err := h.store.AddUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		h.log.Error("failed to add User to storage", slog.Any("user:", user))
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		h.log.Error("failed to response", err.Error())
		return
	}
}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	user, err := h.store.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
