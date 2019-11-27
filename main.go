package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	route.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		_, _ = fmt.Fprintf(w, "Hello %s\n", name)
	})

	_ = http.ListenAndServe(":9990", route)
}
