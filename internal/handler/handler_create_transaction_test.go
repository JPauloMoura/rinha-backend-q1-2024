package handler

import (
	"bytes"
	"io"
	"testing"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

func Benchmark_validateRequest(b *testing.B) {
	js := `{"valor": 100, "tipo": "c", "descricao":"abc"}`

	for n := 0; n < b.N; n++ {
		body := io.NopCloser(bytes.NewReader([]byte(js)))
		var e entities.Transaction
		validateRequest(1, nil, body, &e)
	}
}
