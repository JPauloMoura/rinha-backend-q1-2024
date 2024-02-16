package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/service"
	"github.com/go-chi/chi/v5"
)

func GenerateExtract(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Debug(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	extract, err := service.GenerateExtract(id)
	if err != nil && err.Error() == "client not found" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Debug(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(extract)
}
