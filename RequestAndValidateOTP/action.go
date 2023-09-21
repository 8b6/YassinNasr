package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func requestOTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applciation/json")
	reqBody, mapR := validateInput(r)
	if len(mapR) > 0 {
		var rBody ResponseBody
		rBody.Status = false
		rBody.ResponseCode = "104"
		rBody.ResponseMessage = "please check your parameter"
		responseMessage(w, rBody)
	} else {
		requestOTPcontroller(w, reqBody.PhoneNumber, reqBody.Password, reqBody.Language)
	}

}

func validateOTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applciation/json")
	reqBody, mapR := validateInputV(r)
	if len(mapR) > 0 {
		var rBody ResponseBody
		rBody.Status = false
		rBody.ResponseCode = "104"
		rBody.ResponseMessage = "please check your parameter"
		responseMessage(w, rBody)
	} else {
		validateOTPcontroller(w, reqBody.PhoneNumber, reqBody.OTP)
	}

}

func responseMessage(w http.ResponseWriter, rBody interface{}) {

	writeLog(fmt.Sprintf("responseMessage:    %+v", rBody))

	json.NewEncoder(w).Encode(rBody)

}
