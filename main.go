package main

import (
	"microservice_go/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	handlers.InitHandlers(r)

	err := http.ListenAndServe(":5556", r)

	if err != nil {
		panic(err)
	}
}
