package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// File server for static asset
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Other handlers for the application
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	port := ":4000"
	log.Print("Starting server on ", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
