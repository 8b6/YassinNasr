package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

func insertReservationController(w http.ResponseWriter, insertedReservation Reservation) {

	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	reserveTableID, existence := CheckTableExistence(db, insertedReservation.RestaurantID, insertedReservation.TableType, insertedReservation.ReservetionStartAt)
	if !existence {
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "there is no table available in this time ,try another one  !!"
		responseMessage(w, rBody)
		return
	} else {

		err := InsertReservation(insertedReservation, reserveTableID[0], db)
		if err == nil {
			rBody.Status = true
			rBody.ResponseCode = "4402"
			rBody.ResponseMessage = "your reservation complete "
			responseMessage(w, rBody)
			return
		} else {
			rBody.Status = false
			rBody.ResponseCode = "0"
			rBody.ResponseMessage = "  database error"
			responseMessage(w, rBody)
			return
		}
	}
}

func validateReservationApi(r *http.Request) (Reservation, map[string]interface{}) {
	var req Reservation
	rules := govalidator.MapData{
		"costumerID":         []string{"required"},
		"restaurantID":       []string{"required"},
		"tableType":          []string{"required"},
		"reservetionStartAt": []string{"required"},
		"reservationEndAt":   []string{"required"},
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
