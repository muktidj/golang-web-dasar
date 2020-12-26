package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeIndexGet(write http.ResponseWriter, read *http.Request){
	if read.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
		var err = tmpl.Execute(write, nil)

		if err != nil {
			http.Error(write, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	http.Error(write, "", http.StatusBadRequest)
}

func routeSubmitPost(write http.ResponseWriter, read *http.Request)  {
	if read.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))

		if err := read.ParseForm(); err != nil {
			http.Error(write, err.Error(), http.StatusInternalServerError)
			return
		}

		var name = read.FormValue("name")
		var message = read.FormValue("message")

		var data = map[string]string{"name": name, "message": message}

		if err := tmpl.Execute(write, data); err != nil {
			http.Error(write, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(write, "", http.StatusBadRequest)
}