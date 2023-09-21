package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hi")
	http.HandleFunc("/signIn", addsubscriber)
	port := GetConfigurationParameter("port", "80195")
	fmt.Println("http://localhost:" + port + "/signIn")

	http.ListenAndServe(":"+port, nil)

}
