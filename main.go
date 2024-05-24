package main

import (
	"microservice_go/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handlers.InitHandlers(r)

	err := http.ListenAndServe(":5556", r)

	if err != nil {
		panic(err)
	}
}
