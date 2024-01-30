package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
	fmt.Fprintf(w, "GET request")
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	case http.MethodPost:
		fmt.Fprintf(w, "POST request")
	default:
		fmt.Fprintf(w, "Unsupported  method %s", r.Method)
	}

}
func main() {

	http.HandleFunc("/", handleRequest)

	log.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
