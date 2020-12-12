package main

import (
	"fmt"
	"html/template"
	"net/http"
	pathpkg "path"
)



func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var filepath = pathpkg.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title" : "Learning Golang",
			"name"	: "Ghost Man",
		}

		err = tmpl.Execute(writer, data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

	})


	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	address := "localhost:9000"
	fmt.Println("Server started at %s\n", address)
	//Cara Pertama untuk Start Server

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}


}