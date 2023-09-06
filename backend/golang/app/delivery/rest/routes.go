package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/taimoor99/three-tier-golang/utills"
)

func NewRouter(r *chi.Mux, ctrl UserControllers){

	r.Group(func(rr chi.Router) {
		rr.Use(middleware.Logger)
		rr.Get("/get-user/{id}", ctrl.GetUserByIdHandler)
		rr.Post("/create-user", ctrl.PostCreateUserHandler)
		rr.Get("/users/{limit}/{offset}", ctrl.GetAllUsersHandler)
		rr.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
			utills.WriteJsonRes(writer, 200, nil, "hello world!")
		})
	})

	return
}
