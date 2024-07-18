package route

import (
	"net/http"

	handlers "palm_code_be/src/interface/handlers/health"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func HealthRouter(h handlers.HealthHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", h.Ping)

	return r
}
