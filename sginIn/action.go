package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addsubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applciation/json")
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	SginUpController(reqBody.Msisdn, reqBody.ServiceCode)
	fmt.Fprintf(w, "Subscriber: %+v", reqBody)

}
