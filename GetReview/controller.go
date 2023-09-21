package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

func getReviewController(w http.ResponseWriter, restaurant restaurant) {

	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	restaurant_id, err := getRestaurantID(db, restaurant.RestaurantName)
	if err == nil {
		if restaurant_id != "-1" {

			restReviwes, err := getReview(db, restaurant_id)
			if err == nil {

				responseMessage(w, restReviwes)
				return
			} else {
				rBody.Status = false
				rBody.ResponseCode = "4401"
				rBody.ResponseMessage = "database error : we could not get reviews"
				responseMessage(w, rBody)
				return
			}

		} else {
			rBody.Status = false
			rBody.ResponseCode = "4401"
			rBody.ResponseMessage = "no user with this restaurant name"

			responseMessage(w, rBody)
			return
		}
	} else {
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "database error : we could not get user id "
		responseMessage(w, rBody)
		return
	}

}

func validateGetReviewApi(r *http.Request) (restaurant, map[string]interface{}) {
	var req restaurant
	rules := govalidator.MapData{
		"restaurantName": []string{"required"},
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
