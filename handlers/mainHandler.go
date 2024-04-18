package handlers

import (
	"html/template"
	"net/http"
	errors "asciiArt/errors"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		templates.ExecuteTemplate(w, "input.html", nil)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		templates.ExecuteTemplate(w, "error.html", errors.NotFound)
		return
	}
}
