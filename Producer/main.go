package main

import (
	"microservice_go/configuration"
	"microservice_go/handlers"
	"microservice_go/infrastructure"
	"microservice_go/infrastructure/broker/consumer"
	"microservice_go/infrastructure/broker/producer"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	p := producer.Create(configuration.Configuration)
	c := consumer.MustCreate(configuration.Configuration)

	c.SubscribeTopics([]string{"status_res"}, nil)
	go infrastructure.HandleBrokerMsg(c)

	handlers.InitHandlers(r, p)

	err := http.ListenAndServe(":5555", r)

	if err != nil {
		panic(err)
	}
}
