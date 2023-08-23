package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"yandex/internal/app/repository"

	"github.com/go-chi/chi/v5"
)

func PostAddNewID(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	longURL := string(body)

	_, err = url.ParseRequestURI(longURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := len(repository.InMemoryStorage.InMemory) + 1
	short := "http://127.0.0.1:8080/" + strconv.Itoa(id)

	newItem := repository.Storage{
		ID:    id,
		Long:  longURL,
		Short: short,
	}

	repository.InMemoryStorage.InMemory = append(repository.InMemoryStorage.InMemory, newItem)
	// fmt.Println(repository.InMemoryStorage.InMemory)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(short))
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(repository.InMemoryStorage.InMemory)
	for _, v := range repository.InMemoryStorage.InMemory {
		if id == strconv.Itoa(v.ID) {
			Location := v.Long
			if Location != "" {
				// w.Header().Add("Location", "text")
				w.Header().Set("Location", Location)
				http.Redirect(w, r, Location, http.StatusTemporaryRedirect)

			} else {
				w.WriteHeader(http.StatusBadRequest)
			}

		}

	}

}
