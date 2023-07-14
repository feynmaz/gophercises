package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/feynmaz/gophercises/exercise_3_adventure/config"
	"github.com/feynmaz/gophercises/exercise_3_adventure/handler"
	"github.com/feynmaz/gophercises/exercise_3_adventure/storage"
	"github.com/go-chi/chi/v5"
)

func init() {

}

func main() {
	config := config.GetDefault()

	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	storage, err := storage.NewStorageFromFile(config.StorageFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.PrintContent(); err != nil {
		log.Fatal(err)
	}

	handler := handler.NewHandler(storage, tpl)

	router := chi.NewRouter()
	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	router.Get("/", handler.Index)
	router.Post("/", handler.SubmitStory)

	http.ListenAndServe(":3333", router)
}
