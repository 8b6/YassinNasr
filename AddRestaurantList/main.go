package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func userInformationHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Origin", "*")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
func main() {
	// testing()
	router := mux.NewRouter()
	updateService := router.PathPrefix("/updateService").Subrouter()
	insertService := router.PathPrefix("/insertService").Subrouter()
	updateService.HandleFunc("", authToken(updateServiceHandler)).Methods(http.MethodPost, http.MethodOptions)
	insertService.HandleFunc("", authToken(insertServiceHandler)).Methods(http.MethodPost, http.MethodOptions)

	
	log.Fatal(http.ListenAndServe(":8090", router))
}
