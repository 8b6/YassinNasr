package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	UpdateUserInformation := router.PathPrefix("/downloadImg/{folder}/{pic}").Subrouter()

	UpdateUserInformation.HandleFunc("", downloadImgHandler).Methods(http.MethodPost)

	// freeport, err := GetFreePort()
	// PORT := strconv.Itoa(freeport)
	// fmt.Println(PORT)
	// fmt.Println(err)
	port := os.Getenv("PORT")
	host, _ := os.Hostname()
	var ipv4ToString string
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
			ipv4ToString = ipv4.String()
		}
	}

	println(port)
	println(ipv4ToString + port)
	//log.Fatal(http.ListenAndServe(ipv4ToString+":"+port, router))
	log.Fatal(http.ListenAndServe(":8099", router))

}
