package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

func GetProfilePicController(w http.ResponseWriter, userInfo UserInfo) {
	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	// rBody.IsSend = false
	var err error

	userInfo.UserID, err = getUserIDfromInfo(db, userInfo.PhoneNumber)

	if err == nil {
		if userInfo.UserID == "-1" {
			rBody.Status = false
			rBody.ResponseCode = "167"
			rBody.ResponseMessage = "no such user with this phone nmuber "
			responseMessage(w, rBody)
			return
		} else {
			userPicURL, err1 := getProfilePicURL(db, userInfo.UserID)
			if err1 == nil {
				if userPicURL == "" {
					rBody.Status = true
					rBody.ResponseCode = "1"
					rBody.ResponseMessage = "user dosen't have profile pic :) "
					responseMessage(w, rBody)
					return
				} else {
					profilePic, err2 := readPic(userPicURL)
					if err2 == nil {
						// rBody.Status = true
						// rBody.ResponseCode = "1"
						// rBody.ResponseMessage = "sending the pic"
						// responseMessage(w, rBody)
						w.WriteHeader(http.StatusOK)
						w.Header().Set("Content-Type", "application/octet-stream")
						w.Write(profilePic)
						return
					} else {
						rBody.Status = false
						rBody.ResponseCode = "167"
						rBody.ResponseMessage = "server error : couldn't read the pic"
						responseMessage(w, rBody)
						return
					}
				}
			} else {
				rBody.Status = false
				rBody.ResponseCode = "167"
				rBody.ResponseMessage = "no such user with this phone nmuber "
				responseMessage(w, rBody)
				return
			}
		}
	} else {
		rBody.Status = false
		rBody.ResponseCode = "534"
		rBody.ResponseMessage = "database error: error in getUserIDfromInfo !! "
		responseMessage(w, rBody)
		return
	}

}

func readPic(URL string) ([]byte, error) {
	readedPic, err := ioutil.ReadFile(URL)
	if err != nil {
		return readedPic, err
	}
	return readedPic, err
}

func validateGetProfilePicApi(r *http.Request) (UserInfo, map[string]interface{}) {
	var req UserInfo
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
