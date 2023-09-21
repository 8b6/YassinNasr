package main

import "net/http"

func getMarksController(w http.ResponseWriter) {

	db, _ := SQLConnection()
	//rBody.Ticket = "-1"
	markers, err := GetMarkers(db)

	if err == nil {
		responseMessage(w, markers)
		return
	}

}

