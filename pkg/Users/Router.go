package Users

import (
	"github.com/go-chi/chi"
)

func UserRoute(s Service) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/getUser", getHandler(s))
	r.Post("/updateUser", postHandler)
	return r
}
