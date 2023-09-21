package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	
	http.HandleFunc("/insertReservation", insertReservationHandler)
	// println(port)
	log.Fatal(http.ListenAndServe(":8095", nil))
}


func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "147.135.1.145:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
