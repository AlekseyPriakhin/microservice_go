package handlers

import (
	"microservice_go/repository"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Error struct {
	Message string `json:"message"`
}

func renderError(w http.ResponseWriter, r *http.Request, err string, code int) {
	render.Status(r, code)
	render.JSON(w, r, Error{Message: err})
}

func InitHandlers(r chi.Router) {
	r.Route("/course", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, repository.GetCourse())
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

			render.JSON(w, r, item)
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

			render.JSON(w, r, res)
		})
	})
}
