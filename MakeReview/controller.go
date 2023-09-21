package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
)

func insertReviewontroller(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	//Access the photo key - First Approach
	phone_number := r.FormValue("user_phone_number")
	restaurant_name := r.FormValue("Restaurant_name")
	review := r.FormValue("user_review")
	file, h, err := r.FormFile("review_pic")

	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "server error: could not open the photo"
		responseMessage(w, rBody)
		return
	}
	userID, err := getUserID(db, phone_number)
	if err == nil {
		if userID != "-1" {
			restaurantID, err := getRestaurantID(db, restaurant_name)
			if err == nil {
				if restaurantID != "-1" {
					fileType := GetFileType(h.Filename)
					dir_path := "./"
					reviewPicURL := dir_path + "review_" + userID + "_" + restaurantID + "." + fileType
					tmpfile, err := os.Create(reviewPicURL)
					if err != nil {
						rBody.Status = false
						rBody.ResponseCode = "4401"
						rBody.ResponseMessage = "server error : could not save the photo"
						responseMessage(w, rBody)
						return
					}
					defer tmpfile.Close()
					_, err = io.Copy(tmpfile, file)
					if err != nil {
						rBody.Status = false
						rBody.ResponseCode = "4401"
						rBody.ResponseMessage = "server error : could not save the photo"
						responseMessage(w, rBody)
						return
					}

					if insertUserReview(db, userID, restaurantID, review, reviewPicURL) {
						rBody.Status = true
						rBody.ResponseCode = "1"
						rBody.ResponseMessage = "User information inserted "
						responseMessage(w, rBody)
					} else {
						rBody.Status = false
						rBody.ResponseCode = "4401"
						rBody.ResponseMessage = "database error : could not save the photo url in the database"
						responseMessage(w, rBody)
						return
					}

				} else {
					rBody.Status = false
					rBody.ResponseCode = "4401"
					rBody.ResponseMessage = "no restaurant with this name. "
					responseMessage(w, rBody)
					return
				}
			} else {
				rBody.Status = false
				rBody.ResponseCode = "4401"
				rBody.ResponseMessage = "database error : could not save the photo"
				responseMessage(w, rBody)
				return
			}
		} else {
			rBody.Status = false
			rBody.ResponseCode = "4401"
			rBody.ResponseMessage = "no user with this number. "
			responseMessage(w, rBody)
			return
		}

	} else {
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "database error : could not get the user id "
		responseMessage(w, rBody)
		return
	}
}
