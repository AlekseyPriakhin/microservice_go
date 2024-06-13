package main

import (
	"microservice_go/configuration"
	"microservice_go/handlers"
	"microservice_go/infrastructure"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	p := infrastructure.InitProducerWithAppConfig(configuration.Configuration)

	handlers.InitHandlers(r, p)

	err := http.ListenAndServe(":5555", r)

	if err != nil {
		panic(err)
	}
}
