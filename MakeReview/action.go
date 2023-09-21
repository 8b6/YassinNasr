package main

import (
	"encoding/json"
	"net/http"
)

func insertReviewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "applciation/json")
	
	insertReviewontroller(w, r)
	// }

}

func responseMessage(w http.ResponseWriter, rBody interface{}) {

	//writeLog(fmt.Sprintf("responseMessage:    %+v", rBody))

	json.NewEncoder(w).Encode(rBody)

}
