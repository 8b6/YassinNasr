package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

func UpdateUserInformationController(w http.ResponseWriter, userInfo UserInfo) {
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
			if updateTableUserRegist(db, userInfo) {
				if updateTableUserinfo(db, userInfo) {
					rBody.Status = true
					rBody.ResponseCode = "1"
					rBody.ResponseMessage = "user info updated. "
					responseMessage(w, rBody)
				} else {
					rBody.Status = false
					rBody.ResponseCode = "0"
					rBody.ResponseMessage = "database error : error in update user_information"
					responseMessage(w, rBody)
				}
			} else {
				rBody.Status = false
				rBody.ResponseCode = "167"
				rBody.ResponseMessage = "database error : error in update user_registration"
				responseMessage(w, rBody)
			}
		}

	}

	if err != nil {
		rBody.Status = false
		rBody.ResponseCode = "160"
		rBody.ResponseMessage = " Fail to request OTP"
		responseMessage(w, rBody)
	}

}

func validateUpdateUserInformationApi(r *http.Request) (UserInfo, map[string]interface{}) {
	var req UserInfo
	rules := govalidator.MapData{
		"firstName":   []string{"required"},
		"lastName":    []string{"required"},
		"gender":      []string{"required"},
		"birthdate":   []string{"required"},
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
