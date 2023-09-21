package main

import (
	"database/sql"
	"fmt"
)

func control() {
	var db *sql.DB
	db, _ = SQLConnection()

	var boo bool
	boo = checkUserExistence(db, "11688777")
	fmt.Println("checkUserExistence  ", boo)
}
func SginUpController(msisdn string, phoneNumber string) {
	var db *sql.DB
	db, _ = SQLConnection()
	userID, err := getUserID(db, phoneNumber)
	if err == nil {
		if !checkUserExistence(db, userID) {
			insertUserInformation(db, msisdn, userID)
		}
	}

}
