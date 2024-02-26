package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/go-chi/chi/v5"
)

func (h Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var t entities.Transaction

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	code := validateRequest(
		id, err,
		r.Body,
		&t,
	)

	if code != 200 {
		w.WriteHeader(code)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Minute*2)
	defer cancel()

	resp, err := h.Svc.Repo.CreateTransaction(ctx, id, t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func validateRequest(id int, err error, body io.ReadCloser, e *entities.Transaction) int {
	_, exist := Ids[id]
	if err != nil || !exist {
		return http.StatusNotFound
	}

	b, err := io.ReadAll(body)
	if err != nil {
		return http.StatusUnprocessableEntity
	}

	err = json.Unmarshal(b, e)
	if err != nil {
		return http.StatusUnprocessableEntity
	}

	defer body.Close()

	e.Type = strings.ToLower(e.Type)
	err = e.Validate()
	if err != nil {
		return http.StatusUnprocessableEntity
	}

	return 200
}
