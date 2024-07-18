package route

import (
	"net/http"

	handlers "palm_code_be/src/interface/handlers/media"

	"github.com/go-chi/chi/v5"
)

func MediaRouter(h handlers.MediaHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.Get)
	r.Get("/{id}", h.GetByID)

	return r
}
