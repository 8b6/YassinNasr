package main

import (
	"encoding/json"
	"net/http"
)

func GetProfilePicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-type", "applciation/json")
	reqBody, mapR := validateGetProfilePicApi(r)
	if len(mapR) > 0 {
		var rBody ResponseBody
		rBody.Status = false
		rBody.ResponseCode = "104"
		rBody.ResponseMessage = "please check your parameter"
		responseMessage(w, rBody)
	} else {
		GetProfilePicController(w, reqBody)
	}
}

func responseMessage(w http.ResponseWriter, rBody interface{}) {

	//writeLog(fmt.Sprintf("responseMessage:    %+v", rBody))

	json.NewEncoder(w).Encode(rBody)

}
