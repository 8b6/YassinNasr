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

	insertReview := router.PathPrefix("/insertReview").Subrouter()
	insertReview.HandleFunc("", insertReviewHandler).Methods(http.MethodPost)

	// fmt.Println("http://localhost:" + "8090" + "/recharge")
	// fmt.Println("http://localhost:" + "8090" + "/tokenGeneration")

	log.Fatal(http.ListenAndServe(":8099", router))
}
