package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

func insertServiceController(w http.ResponseWriter, insertedService Service) {

	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody

	if checkServiceExistence(db, insertedService.ServiceCode) {
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "there is a Service with the same service code !!"
		responseMessage(w, rBody)
		return
	} else {
		err := InsertNewService(insertedService, db)
		if err == nil {
			rBody.Status = true
			rBody.ResponseCode = "4402"
			rBody.ResponseMessage = "Service added successfully"
			responseMessage(w, rBody)
			return
		} else {
			rBody.Status = false
			rBody.ResponseCode = "0"
			rBody.ResponseMessage = "  internal error from Sudatel API"
			responseMessage(w, rBody)
			return
		}
	}
}

func updateServiceController(w http.ResponseWriter, insertedService UpdateService) {

	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody

	if !checkServiceExistence(db, insertedService.ServiceCode) {
		rBody.Status = false
		rBody.ResponseCode = "4404"
		rBody.ResponseMessage = "this service not exist !!"
		responseMessage(w, rBody)
		return
	}
	serivceId, err := getServiceID(db, insertedService.ServiceCode)
	if err != nil {
		rBody.Status = false
		rBody.ResponseCode = "4409"
		rBody.ResponseMessage = "error can't get service id"
		responseMessage(w, rBody)
		return
	} else {
		err := updateService(*db, insertedService, serivceId)
		if err != nil {
			rBody.Status = false
			rBody.ResponseCode = "4449"
			rBody.ResponseMessage = "example error"
			responseMessage(w, rBody)
			return
		} else {
			rBody.Status = true
			rBody.ResponseCode = "1"
			rBody.ResponseMessage = " update done"
			responseMessage(w, rBody)
		}
	}

}

func validateInputInsertServiceApi(r *http.Request) (Service, map[string]interface{}) {
	var req Service
	rules := govalidator.MapData{
		"serviceCode":  []string{"required"},
		"servicePrice": []string{"required"},
		"ServiceType":  []string{"required"},
		"description":  []string{"required"},
		"homePage":     []string{"required"},
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

func validateInputUpdateServiceApi(r *http.Request) (UpdateService, map[string]interface{}) {
	var req UpdateService
	rules := govalidator.MapData{
		"serviceCode": []string{"required"},
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
