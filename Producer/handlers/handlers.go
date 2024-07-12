package handlers

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"microservice_go/application"
	"microservice_go/repository"
	"microservice_go/utils/mapper"
	"net/http"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Error struct {
	Message string `json:"message"`
}

func renderError(w http.ResponseWriter, r *http.Request, err string, code int) {
	render.Status(r, code)
	render.JSON(w, r, Error{Message: err})
}

func InitHandlers(r chi.Router, p *kafka.Producer) {
	m := application.GetRequestMap()

	r.Route("/course", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, mapper.ToCourseResDtoList(repository.GetCourse()))
		})

		r.Get("/request", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, m)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {

			id, idErr := strconv.Atoi(chi.URLParam(r, "id"))
			if idErr != nil {
				renderError(w, r, idErr.Error(), http.StatusBadRequest)
				return
			}

			item, err := repository.FindCourse(id)
			if err != nil {
				renderError(w, r, err.Error(), http.StatusNotFound)
				return
			}

			topic := "status_req"
			guid := uuid.New().String()
			message, _ := json.Marshal(application.StatusChangeReqBrokerMsg{
				Course: item,
				ReqId:  guid,
			})

			users := repository.Users

			rand, _ := rand.Int(rand.Reader, big.NewInt(int64(len(users))))
			application.AddRequestToMap(guid, users[rand.Int64()])

			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
				Value:          []byte(message),
			}, nil)

			if err != nil {
				renderError(w, r, err.Error(), http.StatusInternalServerError)
				return
			}

			render.JSON(w, r, mapper.ToCourseResDto(item))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			req := repository.CourseReqDto{}
			err := render.DecodeJSON(r.Body, &req)
			if err != nil {
				renderError(w, r, "bad request", http.StatusBadRequest)
				return
			}

			res, addErr := repository.AddCourse(req)

			if addErr != nil {
				renderError(w, r, addErr.Error(), http.StatusBadRequest)
				return
			}

			render.Status(r, http.StatusCreated)
			render.JSON(w, r, mapper.ToCourseResDto(res))
		})
	})
}
