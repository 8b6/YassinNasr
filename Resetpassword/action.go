package main

import (
	"encoding/json"
	"net/http"
)

func resetPassword_requestOTP_Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "applciation/json")
	reqBody, mapR := validateResetPassword_requestOTPApi(r)
	if len(mapR) > 0 {
		var rBody ResponseBody
		rBody.Status = false
		rBody.ResponseCode = "104"
		rBody.ResponseMessage = "please check your parameter"
		responseMessage(w, rBody)
	} else {
		resetPassword_requestOTP_Controller(w, reqBody)
	}

}

func resetPassword_setNewPassword_Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "applciation/json")
	reqBody, mapR := validateresetPassword_setNewPasswordApi(r)
	if len(mapR) > 0 {
		var rBody ResponseBody
		rBody.Status = false
		rBody.ResponseCode = "104"
		rBody.ResponseMessage = "please check your parameter"
		responseMessage(w, rBody)
	} else {
		resetPassword_setNewPasswordController(w, reqBody)
	}

}

func responseMessage(w http.ResponseWriter, rBody interface{}) {

	//writeLog(fmt.Sprintf("responseMessage:    %+v", rBody))

	json.NewEncoder(w).Encode(rBody)

}
