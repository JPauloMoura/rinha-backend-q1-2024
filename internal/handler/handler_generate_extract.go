package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h Handler) GenerateExtract(w http.ResponseWriter, r *http.Request) {
	// defer timeTrack(time.Now(), "GenerateExtract")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	_, exist := Ids[id]
	if err != nil || !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	extract, err := h.Svc.GenerateExtract(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(extract)
}
