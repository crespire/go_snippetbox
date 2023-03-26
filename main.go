package main

import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello from Snippetbox"))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)

  port := ":4000"
  log.Print("Starting server on ", port)
  err := http.ListenAndServe(port, mux)
  log.Fatal(err)
}
