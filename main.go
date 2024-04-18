package main

import (
	"asciiArt/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
