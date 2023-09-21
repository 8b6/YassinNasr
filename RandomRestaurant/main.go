package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	//testing()
	router := mux.NewRouter()
	randomRestaurant := router.PathPrefix("/randomRestaurant").Subrouter()
	randomRestaurant.HandleFunc("", randomRestaurantHandler).Methods(http.MethodGet)

	// fmt.Println("http://localhost:" + "8090" + "/recharge")
	// fmt.Println("http://localhost:" + "8090" + "/tokenGeneration")

	log.Fatal(http.ListenAndServe(":8099", router))
}
