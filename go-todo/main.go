package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Item string
	Done bool
}

var tmpl *template.Template

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn Go", Done: false},
		},
	}
	tmpl.Execute(w, data)
}
func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("template/template.gohtml"))
	mux.HandleFunc("/todo", todo)

	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/",http.StripPrefix("/static/",fs))
	log.Fatal(http.ListenAndServe(":9091", mux))
	
}
