package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/thedevsaddam/govalidator"
)

func resetPassword_setNewPasswordController(w http.ResponseWriter, userInfo RequestBody) {
	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	userID, err := getUserID(db, userInfo.PhoneNumber)
	if err == nil {
		if userID != "-1" {
			err := UpdatePassword(db, userID, userInfo.Password)
			if err == nil {
				rBody.Status = true
				rBody.ResponseCode = "1"
				//rBody.IsSend = true
				rBody.ResponseMessage = "Password updated"
				responseMessage(w, rBody)
			} else {
				rBody.Status = false
				rBody.ResponseCode = "0"
				//rBody.IsSend = true
				rBody.ResponseMessage = "database error : error in update password in user_registration table!!!"
				responseMessage(w, rBody)
			}
		} else {
			rBody.Status = false
			rBody.ResponseCode = "0"
			//rBody.IsSend = true
			rBody.ResponseMessage = "no user id for the number"
			responseMessage(w, rBody)
		}
	} else {
		rBody.Status = false
		rBody.ResponseCode = "0"
		//rBody.IsSend = true
		rBody.ResponseMessage = "database error : error in get user id from user_registration table!!!"
		responseMessage(w, rBody)
	}

}

func resetPassword_requestOTP_Controller(w http.ResponseWriter, userInfo RequestBody) {
	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	// rBody.IsSend = false
	user_id, err := getUserIDfromInfo(db, userInfo.PhoneNumber)

	if err == nil {
		if user_id == "-1" {
			rBody.Status = false
			rBody.ResponseCode = "167"
			rBody.ResponseMessage = "no such user with this phone nmuber "
			responseMessage(w, rBody)
			return
		}
		actualRequestOTP(w, db, userInfo)
	}

	if err != nil {
		rBody.Status = false
		rBody.ResponseCode = "160"
		rBody.ResponseMessage = " Fail to request OTP"
		responseMessage(w, rBody)
	}

}

func actualRequestOTP(w http.ResponseWriter, db *sql.DB, userinfo RequestBody) {
	var rBody ResponseBody
	pinCode := strconv.Itoa(randomInt(100000, 999999))
	err := insertOTP_request(db, userinfo.PhoneNumber, userinfo.Password, pinCode)
	if err == nil {
		rBody.Status = true
		rBody.ResponseCode = "1"
		//rBody.IsSend = true
		rBody.ResponseMessage = "The OTP sent Successfully"
		responseMessage(w, rBody)
	} else {
		rBody.Status = false
		rBody.ResponseCode = "0"
		//rBody.IsSend = true
		rBody.ResponseMessage = "  database insertion error"
		responseMessage(w, rBody)
	}

}

func validateresetPassword_setNewPasswordApi(r *http.Request) (RequestBody, map[string]interface{}) {
	var req RequestBody
	rules := govalidator.MapData{
		"phoneNumber": []string{"required"},
		"password":    []string{"required"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &req,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	fmt.Println(req) // your incoming JSON data in Go data struct
	err := map[string]interface{}{"validationError": e}
	var m url.Values
	x := fmt.Sprintf("%v", e)
	y := fmt.Sprintf("%v", m)
	if x != y {
		return req, err
	}
	return req, nil
}

func validateResetPassword_requestOTPApi(r *http.Request) (RequestBody, map[string]interface{}) {
	var req RequestBody
	rules := govalidator.MapData{
		"phoneNumber": []string{"required"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &req,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	fmt.Println(req) // your incoming JSON data in Go data struct
	err := map[string]interface{}{"validationError": e}
	var m url.Values
	x := fmt.Sprintf("%v", e)
	y := fmt.Sprintf("%v", m)
	if x != y {
		return req, err
	}
	return req, nil
}
