package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func insertServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "applciation/json")
	reqBody, mapR := validateInputInsertServiceApi(r)
	if len(mapR) > 0 {
		var rBody ResponseBody
		rBody.Status = false
		rBody.ResponseCode = "104"
		rBody.ResponseMessage = "please check your parameter"
		responseMessage(w, rBody)
	} else {
		insertServiceController(w, reqBody)
	}

}

func updateServiceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "applciation/json")
	reqBody, mapR := validateInputUpdateServiceApi(r)
	if len(mapR) > 0 {
		var rBody ResponseBody
		rBody.Status = false
		rBody.ResponseCode = "104"
		rBody.ResponseMessage = "please check your parameter"
		responseMessage(w, rBody)
	} else {
		updateServiceController(w, reqBody)
	}
}

func responseMessage(w http.ResponseWriter, rBody interface{}) {

	writeLog(fmt.Sprintf("responseMessage:    %+v", rBody))

	json.NewEncoder(w).Encode(rBody)

}
