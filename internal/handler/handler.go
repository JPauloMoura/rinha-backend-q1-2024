package handler

import (
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/service"
)

type Handler struct {
	Svc service.Service
}

var Ids = map[int]bool{
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
}
