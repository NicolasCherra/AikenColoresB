package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var local = "3000"
var produccion = os.Getenv("PORT")
var port = ""

// No estara cargnado el modulo ?
func main() {

	if produccion != "" {
		port = produccion
	} else {
		port = local
	}

	ConnectMongoDB()

	router := mux.NewRouter()
	router.HandleFunc("/", Inicio).Methods("GET")
	router.HandleFunc("/souvenir/{_id}", UpdateSouvenir).Methods("PUT")
	router.HandleFunc("/souvenir", GetSouvenirs).Methods("GET")
	router.HandleFunc("/souvenir/{_id}", GetSouvenir).Methods("GET")
	router.HandleFunc("/souvenir", CreateSouvenir).Methods("POST")
	router.HandleFunc("/souvenir/{_id}", DeleteSouvenir).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
