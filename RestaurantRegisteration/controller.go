package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

func registController(w http.ResponseWriter, registinformation RestaurantRegistration) {

	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	if !CheckRestaurantExistence(db, registinformation.Email) {
		err := insertUserRegistInfo(db, registinformation)
		if err == nil {
			if sendEmail(registinformation.Email) == "" {
				rBody.Status = true
				rBody.ResponseCode = "1"
				rBody.ResponseMessage = "restaurant sign up complete"
				responseMessage(w, rBody)
			} else {
				rBody.Status = false
				rBody.ResponseCode = "0"
				rBody.ResponseMessage = "User registration completed but faild in sending verification email"
				responseMessage(w, rBody)
				return
			}
		} else {
			rBody.Status = false
			rBody.ResponseCode = "4401"
			rBody.ResponseMessage = "database error"
			responseMessage(w, rBody)
			return
		}

	} else {
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "this phone number already use for existed user"
		responseMessage(w, rBody)
		return
	}
}

func validateregistApi(r *http.Request) (RestaurantRegistration, map[string]interface{}) {
	var req RestaurantRegistration
	rules := govalidator.MapData{
		"email":    []string{"required"},
		"password": []string{"required"},
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
