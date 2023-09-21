package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/thedevsaddam/govalidator"
)

func uploadUserImgController(w http.ResponseWriter, r *http.Request) {
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
	file, h, err := r.FormFile("profile_pic")

	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "server error: could not open the photo"
		responseMessage(w, rBody)
		return
	}
	userID, err := getUserIDfromInfo(db, phone_number)
	if err == nil {
		if userID != "-1" {
			fileType := GetFileType(h.Filename)
			dir_path := "./"
			profilePicURL := dir_path + userID + "." + fileType
			tmpfile, err := os.Create(profilePicURL)

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

			if insertUserPicURL(db, userID, profilePicURL) {
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
			rBody.ResponseMessage = "no user with this phone number"
			responseMessage(w, rBody)
			return
		}
	} else {
		rBody.Status = false
		rBody.ResponseCode = "4401"
		rBody.ResponseMessage = "database error : we could not get user id"
		responseMessage(w, rBody)
		return
	}

}

func insertUserInfoController(w http.ResponseWriter, insertedUserInfo UserInfo) {

	var db *sql.DB
	db, _ = SQLConnection()
	defer db.Close()
	var rBody ResponseBody
	UserID, err := getUserID(db, insertedUserInfo.PhoneNumber)
	if err == nil {
		if UserID != "-1" {
			insertedUserInfo.UserID = UserID
			err := insertUserInfo(db, insertedUserInfo)
			if err == nil {
				rBody.Status = true
				rBody.ResponseCode = "1"
				rBody.ResponseMessage = "User information inserted "
				responseMessage(w, rBody)
				return
			} else {
				rBody.Status = false
				rBody.ResponseCode = "4401"
				rBody.ResponseMessage = "database error : we could not insert user information"
				responseMessage(w, rBody)
				return
			}

		} else {
			rBody.Status = false
			rBody.ResponseCode = "4401"
			rBody.ResponseMessage = "no user with this phone number"

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

func validateinsertUserInfoApi(r *http.Request) (UserInfo, map[string]interface{}) {
	var req UserInfo
	rules := govalidator.MapData{
		"firstName":   []string{"required"},
		"lastName":    []string{"required"},
		"gender":      []string{"required"},
		"birthdate":   []string{"required"},
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

func uploadUserImgHandlerApi(w http.ResponseWriter, r *http.Request) (err1 string) {
	// function body of a http.HandlerFunc
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20)
	reader, err := r.MultipartReader()
	if err != nil {
		err1 = "bad request"
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return err1
	}
	// parse text field
	text := make([]byte, 512)
	p, err := reader.NextPart()
	// one more field to parse, EOF is considered as failure here
	if err != nil {
		err1 = "the number of field is not correct "
		return err1
	}
	if p.FormName() != "user_phone_number" {
		// http.Error(w, "text_field is expected", http.StatusBadRequest)
		err1 = "expected field with another name"
		return err1
	}
	_, err = p.Read(text)
	if err != nil && err != io.EOF {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		err1 = "server error"
		return
	}
	// parse file field
	p, err = reader.NextPart()
	if err != nil && err != io.EOF {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		err1 = "the number of field is not correct"
		return
	}
	if p.FormName() != "profile_pic" {
		// http.Error(w, "file_field is expected", http.StatusBadRequest)
		err1 = "expected field with another name"
		return err1
	}
	buf := bufio.NewReader(p)
	
	f, err := ioutil.TempFile("", "")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		err1 = "server error"
		return err1
	}
	defer f.Close()
	var maxSize int64 = 32 << 20
	lmt := io.MultiReader(buf, io.LimitReader(p, maxSize-511))
	written, err := io.Copy(f, lmt)
	if err != nil && err != io.EOF {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		err1 = "server error"
		return err1
	}
	if written > maxSize {
		os.Remove(f.Name())
		// http.Error(w, "file size over limit", http.StatusBadRequest)
		err1 = "file size over limit"
		return err1
	}
	err1 = ""
	// schedule for other stuffs (s3, scanning, etc.)
	return err1
}
