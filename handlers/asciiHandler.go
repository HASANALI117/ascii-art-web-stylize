package handlers

import (
	asciiArt "asciiArt/asciiArt"
	errors "asciiArt/errors"
	"net/http"
	"strings"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		str := r.FormValue("input")
		font := r.FormValue("font")

		switch strings.ToLower(font) {
		case "standard", "shadow", "thinkertoy":
			// font is valid, do nothing
		default:
			// http.Error(w, "Invalid font option", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			templates.ExecuteTemplate(w, "error.html", errors.BadRequest("Invalid font option"))
			return
		}

		lines := strings.Split(str, "\r\n")

		var asciiLines []string
		for _, line := range lines {
			if line == "" {
				continue
			}
			asciiLine, err := asciiArt.AsciiLine(line, font)
			if err != nil {
				// http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
				templates.ExecuteTemplate(w, "error.html", errors.BadRequest(err.Error()))
				return
			}
			asciiLines = append(asciiLines, asciiLine...)
		}

		art := asciiArt.PrintAscii(asciiLines)

		err := templates.ExecuteTemplate(w, "result.html", art)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			templates.ExecuteTemplate(w, "error.html", errors.InternalServer(err.Error()))
			return
		}

	default:
		// http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		templates.ExecuteTemplate(w, "error.html", errors.MethodNotAllowed)
	}

}
