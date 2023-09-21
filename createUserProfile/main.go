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
	insertUserInfo := router.PathPrefix("/insertUserInfo").Subrouter()
	uploadUserImg := router.PathPrefix("/uploadUserImg").Subrouter()
	insertUserInfo.HandleFunc("", insertUserInfoHandler).Methods(http.MethodPost)
	uploadUserImg.HandleFunc("", uploadUserImgHandler).Methods(http.MethodPost)

	// fmt.Println("http://localhost:" + "8090" + "/recharge")
	// fmt.Println("http://localhost:" + "8090" + "/tokenGeneration")

	log.Fatal(http.ListenAndServe(":8099", router))
}
