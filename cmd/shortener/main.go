package main

import (
	"net/http"
	"yandex/internal/app/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Post("/", handlers.PostAddNewID)
	r.Get("/{id}", handlers.GetByID)

	http.ListenAndServe(":8080", r)
}
