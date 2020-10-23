package Router

import (
	"fmt"
	"go_chi/pkg/Database"
	"go_chi/pkg/Users"
	"net/http"

	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {
	r, err := Database.SetUpStorage()

	if err != nil {
		fmt.Println(err)
	}
	us := Users.NewService(r)
	router := chi.NewRouter()
	router.Mount("/api/users", Users.UserRoute(us))
	http.ListenAndServe(":3000", router)

	return router
}
