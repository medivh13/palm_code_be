package route

import (
	"net/http"

	handlers "palm_code_be/src/interface/handlers/page"

	"github.com/go-chi/chi/v5"
)

func PageRouter(h handlers.PageHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Create)
	r.Get("/", h.Get)
	r.Get("/{id}", h.GetByID)
	r.Put("/", h.Update)
	r.Delete("/", h.Delete)

	return r
}
