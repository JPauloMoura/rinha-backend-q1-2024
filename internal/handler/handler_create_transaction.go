package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/service"
	errors "github.com/JPauloMoura/rinha-backend-q1-2024/pkg/errors"
	"github.com/go-chi/chi/v5"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	_, exist := Ids[id]
	if err != nil || !exist {
		// slog.Debug("id not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var t entities.Transaction
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		// slog.Debug(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	t.Type = strings.ToLower(t.Type)
	err = t.Validate()
	if err != nil {
		// slog.Debug(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	resp, err := service.CreateTransaction(context.TODO(), id, t)
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
	json.NewEncoder(w).Encode(resp)
}
