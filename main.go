package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
			case "POST":
				writer.Write([]byte("post"))

			case "GET":
				writer.Write([]byte("get"))

			default:
				http.Error(writer, "", http.StatusBadRequest )


		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}