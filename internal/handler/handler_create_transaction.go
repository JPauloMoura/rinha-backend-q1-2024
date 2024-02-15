package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/service"
	"github.com/go-chi/chi/v5"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Debug(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var t entities.Transaction
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		slog.Debug(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.Type = strings.ToLower(t.Type)
	err = t.Validate()
	if err != nil {
		slog.Debug(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = service.CreateTransaction(id, t)
	if err != nil && err.Error() == "client not found" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil && err.Error() == "transaction invalid" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
