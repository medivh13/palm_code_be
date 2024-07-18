package route

import (
	"net/http"

	handlers "palm_code_be/src/interface/handlers/upload"

	"github.com/go-chi/chi/v5"
)

func UploadRouter(h handlers.UploadHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Upload)

	return r
}
