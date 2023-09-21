package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//testing()
	fmt.Println("Hi")
	router := mux.NewRouter()
	requestOTP := router.PathPrefix("/requestOTP").Subrouter()
	requestOTP.HandleFunc("", authToken(requestOTPHandler)).Methods(http.MethodPost)
	validateOTP := router.PathPrefix("/validateOTP").Subrouter()
	validateOTP.HandleFunc("", authToken(validateOTPHandler)).Methods(http.MethodPost)
	port := GetConfigurationParameter("port", "8015")
	fmt.Println("http://localhost:" + port + "/requestOTP")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
