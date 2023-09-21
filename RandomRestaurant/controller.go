package main

import (
	"database/sql"
	"net/http"
)

func randomRestaurantController(w http.ResponseWriter) {
	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	// rBody.IsSend = false
	var err error
	restaurantID, err := AllRestaurantIDs(db)
	if err == nil {
		randomRest := RandomId(restaurantID, 4)
		restaurants, err := Getrandrest(db, randomRest)
		if err == nil {
			responseMessage(w, restaurants)
			return
		} else {
			rBody.Status = false
			rBody.ResponseCode = "167"
			rBody.ResponseMessage = "database error: error in Getrandrest "
			responseMessage(w, rBody)
		}
	} else {
		rBody.Status = false
		rBody.ResponseCode = "167"
		rBody.ResponseMessage = "database error: error in AllRestaurantIDs"
		responseMessage(w, rBody)
	}

}
