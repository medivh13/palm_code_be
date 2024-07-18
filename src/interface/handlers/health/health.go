package handlers

import (
	"net/http"

	"palm_code_be/src/interface/response"
)

type HealthHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
	response response.IResponseClient
}

func NewHealthHandler(r response.IResponseClient) HealthHandler {
	return &healthHandler{
		response: r,
	}
}

func (h *healthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	h.response.JSON(w, "Pong", nil, nil)
}
