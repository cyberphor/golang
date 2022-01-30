package main

import (
	"net/http"

	"github.com/cyberphor/demo/controllers"
)

func main() {
	filePath := http.Dir("./views/")
	fileServer := http.FileServer(filePath)
	http.Handle("/static/", fileServer)
	http.HandleFunc("/", controllers.LoginPage)
	http.HandleFunc("/login", controllers.Login)
	http.ListenAndServe(":8000", nil)
}
