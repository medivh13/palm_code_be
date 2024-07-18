package route

import (
	"net/http"

	handlers "palm_code_be/src/interface/handlers/user"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h handlers.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)

	return r
}
