package handler

import (
	"context"
	"encoding/json"

	// "log/slog"
	"net/http"
	"strconv"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/service"
	errors "github.com/JPauloMoura/rinha-backend-q1-2024/pkg/errors"
	"github.com/go-chi/chi/v5"
)

var Ids = map[int]bool{
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
}

func GenerateExtract(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	_, exist := Ids[id]
	if err != nil || !exist {
		// slog.Debug("id not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	extract, err := service.GenerateExtract(context.TODO(), id)
	if err == errors.ErrClientNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(extract)
}
