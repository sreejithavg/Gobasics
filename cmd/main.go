package main

import (
	"firstProject/pkg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/home", pkg.HomeLink)
	err := http.ListenAndServe(":8080", router)
	log.Println("server localhost:8080 started")
	log.Fatal(err)
}
