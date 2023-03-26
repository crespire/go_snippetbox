package main

import (
  "log"
  "fmt"
  "net/http"
  "strconv"
)
func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  _, err := w.Write([]byte("Hello from Snippetbox"))
  if err != nil {
    log.Print(err)
  }
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  fmt.Fprintf(w, "Display snippet %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }

  _, err := w.Write([]byte("Create a new snippet..."))
  if err != nil {
    log.Print(err)
  }
}

