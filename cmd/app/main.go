package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello World"))
}

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)

	if err := http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
