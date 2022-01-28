package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":8888", nil)
}
