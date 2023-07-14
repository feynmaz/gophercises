package handler

import (
	"html/template"
	"net/http"

	"github.com/feynmaz/gophercises/exercise_3_adventure/storage"
)

type handler struct {
	storage *storage.Storage
	tpl     *template.Template
}

func NewHandler(storage *storage.Storage, tpl *template.Template) *handler {
	return &handler{
		storage: storage,
		tpl:     tpl,
	}
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	initialStory, err := h.storage.GetInitial()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	h.tpl.ExecuteTemplate(w, "index.html", initialStory)
}

func (h *handler) SubmitStory(w http.ResponseWriter, r *http.Request) {
	storyID := r.FormValue("arc")
	story, err := h.storage.GetById(storyID)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	h.tpl.ExecuteTemplate(w, "index.html", story)
}
