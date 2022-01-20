package main

import (
    "fmt"
    "net/http"
)

func ctfConsole(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", "foo")
}

func main() {
    http.HandleFunc("/", ctfConsole)
    http.ListenAndServe(":5050", nil)
}
