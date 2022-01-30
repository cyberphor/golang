package main

import (
	"html/template"
)

/* demo1
func main() {
	http.HandleFunc("/", Browse)
	http.HandleFunc("/index.html", Index)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":8888", nil)
}
*/

type Page struct {
	Title template.HTML
	Body  template.HTML
}

func main() {
	t, err := template.ParseFiles("template.gohtml")
	if err != nil {
		panic(err)
	}

	page := Page{
		Title: "Welcome",
		Body:  template.HTML(iouti.ReadFile("")),
	}

	err
}
