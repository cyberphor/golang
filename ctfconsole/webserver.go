package main

import (
	"net/http"
)

func ServeHtml() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.ListenAndServe(":8000", nil)
}