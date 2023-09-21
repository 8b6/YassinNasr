package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/thedevsaddam/govalidator"
)

func requestOTPcontroller(w http.ResponseWriter, phoneNumber string, password string, language string) {
	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	rBody.IsSend = false
	user_id, err := getUserID(db, phoneNumber)

	if err == nil {
		if !checkUserExistence(db, user_id) {
			actualRequestOTP(w, db, phoneNumber, password, language)
		} else {
			rBody.Status = false
			rBody.ResponseCode = "111"
			rBody.ResponseMessage = "Your phone number already subscribed in this app."
			responseMessage(w, rBody)
			return
		}
	}

	if err != nil {
		rBody.Status = false
		rBody.ResponseCode = "160"
		rBody.ResponseMessage = " Fail to request OTP"
		responseMessage(w, rBody)
	}

}

func validateOTPcontroller(w http.ResponseWriter, phoneNumber string, OTP string) {

	db, _ := SQLConnection()
	var rBody ValidResponseBody
	rBody.Ticket = "-1"
	serviceID, err := getUserID(db, phoneNumber)

	if err == nil {
		if serviceID == "-1" {
			rBody.Status = false
			rBody.ResponseCode = "166"
			rBody.ResponseMessage = "Service not  found"
			responseMessage(w, rBody)
			return
		}
		if !checkUserExistence(db, phoneNumber) {
			actualValidateOTP(w, db, phoneNumber, OTP)
		} else {
			rBody.Status = false
			rBody.ResponseCode = "111"
			rBody.ResponseMessage = "Your MSISDN already subscribed in this service."
			responseMessage(w, rBody)
			return
		}
	}

	if err != nil {
		rBody.Status = false
		rBody.ResponseCode = "162"
		rBody.ResponseMessage = " Fail to validate OTP"
		responseMessage(w, rBody)
	}

}

func actualValidateOTP(w http.ResponseWriter, db *sql.DB, phoneNumber string, OTP string) {
	var rBody ValidResponseBody
	checkOTP, resp := detialValidateOTP(db, phoneNumber, OTP)

	if checkOTP {
		rBody.Status = true
		rBody.ResponseCode = "1"
		rBody.Ticket = OTP
		rBody.ResponseMessage = "The OTP is  correct"
		updateStatusOfOTP(db, phoneNumber)
		responseMessage(w, rBody)

	} else {
		rBody = resp
		responseMessage(w, rBody)
	}
}

func detialValidateOTP(conn *sql.DB, phoneNumber string, OTP string) (boo bool, rBody ValidResponseBody) {
	if checkOTPInvalid(conn, phoneNumber, OTP) {
		if checkOTPused(conn, phoneNumber, OTP) {
			if checkOTPexpird(conn, phoneNumber, OTP) {
				return true, rBody
			} else {
				rBody.Status = false
				rBody.ResponseCode = "165"
				rBody.Ticket = "-1"
				rBody.ResponseMessage = " The pin expird "
				return false, rBody
			}
		} else {
			rBody.Status = false
			rBody.ResponseCode = "164"
			rBody.Ticket = "-1"
			rBody.ResponseMessage = " The pin Used "
			return false, rBody
		}

	} else {
		rBody.Status = false
		rBody.ResponseCode = "163"
		rBody.Ticket = "-1"
		rBody.ResponseMessage = " The pin invalid "
		return false, rBody
	}

}

func actualRequestOTP(w http.ResponseWriter, db *sql.DB, phoneNumber string, password string, language string) {
	var rBody ResponseBody
	pinCode := strconv.Itoa(randomInt(100000, 999999))
	if SendSMS(phoneNumber, pinCode, language) {
		err := insertOTP_request(db, phoneNumber, password, pinCode)
		if err == nil {
			rBody.Status = true
			rBody.ResponseCode = "1"
			rBody.IsSend = true
			rBody.ResponseMessage = "The OTP sent Successfully"
			responseMessage(w, rBody)
		} else {
			rBody.Status = false
			rBody.ResponseCode = "0"
			rBody.IsSend = true
			rBody.ResponseMessage = "  internal error from Sudatel API"
			responseMessage(w, rBody)
		}
	} else {
		rBody.Status = false
		rBody.ResponseCode = "161"
		rBody.IsSend = false
		rBody.ResponseMessage = " Fail to sent  SMS"
		responseMessage(w, rBody)
	}

}

func validateInput(r *http.Request) (RequestBody, map[string]interface{}) {
	var req RequestBody
	rules := govalidator.MapData{
		"msisdn":      []string{"required", "numeric", "between:12,12"},
		"serviceCode": []string{"required"},
		"language":    []string{"required"},
		"price":       []string{"required"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &req,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	err := map[string]interface{}{"validationError": e}
	var m url.Values
	x := fmt.Sprintf("%v", e)
	y := fmt.Sprintf("%v", m)
	if x != y {
		return req, err
	}
	return req, nil
}

func validateInputV(r *http.Request) (RequestBody, map[string]interface{}) {
	var req RequestBody
	rules := govalidator.MapData{
		"msisdn":      []string{"required", "numeric", "between:12,12"},
		"serviceCode": []string{"required"},
		"pinCode":     []string{"required"},
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
