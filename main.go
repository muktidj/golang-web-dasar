package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"io"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeIndexGet(write http.ResponseWriter, read *http.Request){
	if read.Method != "GET" {
		http.Error(write, "", http.StatusBadRequest)
		return
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		var err = tmpl.Execute(write, nil)

		if err != nil {

			http.Error(write, err.Error(), http.StatusInternalServerError)

	}
}

func routeSubmitPost(write http.ResponseWriter, read *http.Request)  {
	if read.Method != "POST" {
		http.Error(write, "", http.StatusBadRequest)
		return
	}

	if err := read.ParseMultipartForm(1024); err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := read.FormValue("alias")

	uploadedFile, handler, err := read.FormFile("file")
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()


	dir, err := os.Getwd()
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}

	write.Write([]byte("done"))

}