package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var databaseServer, databaseUser, database, password string


func SQLConnection() (result *sql.DB, err error) {

	connectionString := fmt.Sprintf("%v:%v@tcp(%s:3306)/%v?parseTime=true",
		databaseUser, password, databaseServer, database)
	var db *sql.DB
	db, err = sql.Open("mysql", connectionString)
	if err != nil {

		writeLog("Error in SQLConnection" + err.Error())
		fmt.Print("Error connecting to: ", databaseServer, ": ", err.Error())

	} else {

		fmt.Println("Connected to: ", databaseServer)

	}

	return db, err
}

func getUserID(connection *sql.DB, user_phone_number string) (id string, err error) {

	sqlStatement := `SELECT user_id FROM user_registration where user_phone=` + user_phone_number + `; `
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	var user_id string
	user_id = "-1"
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&user_id)

		}

	} else {
		writeLog("Error in getUserID: " + err.Error())
	}
	fmt.Println(err)

	return user_id, err

}

func checkUserExistence(connection *sql.DB, userID string) bool {
	sqlStatement := `SELECT * FROM user_registration where user_id=` + userID
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	count := 0
	if err == nil {
		for rows.Next() {

			count++
		}

	} else {
		writeLog("Error in checkUserExistence: " + err.Error())

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}

func insertOTP_request(connection *sql.DB, user_phone_number string, userPassword string, OTP string) (err error) {

	sql := `INSERT INTO nowru.user_registration_OTP (phone_num,password,otp) VALUES (?,?,?);
	`

	_, err = connection.Exec(sql, user_phone_number, userPassword, OTP)

	if err != nil {

		writeLog("Error in insertOTP_request: " + err.Error())

	}
	return
}

func checkOTPexistence(connection *sql.DB, user_phone_number string, OTP string) bool {
	sqlStatement := `SELECT * FROM user_registration_OTP where phone_num=` + user_phone_number + ` and otp = ` + OTP + ` and  status=0 and TIMESTAMPDIFF(HOUR,creation_date,CURRENT_TIMESTAMP())< 1 `
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	count := 0
	if err == nil {
		for rows.Next() {

			count++
		}

	} else {
		writeLog("Error in checkOTPexistence: " + err.Error())

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}
func checkOTPInvalid(connection *sql.DB, phoneNumber string, OTP string) bool {
	sqlStatement := `SELECT * FROM user_registration_OTP where phone_num=` + phoneNumber + ` and otp = ` + OTP
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	count := 0
	if err == nil {
		for rows.Next() {

			count++
		}

	} else {
		writeLog("Error in checkOTPInvalid: " + err.Error())

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}
func checkOTPused(connection *sql.DB, phoneNumber string, OTP string) bool {
	sqlStatement := `SELECT * FROM user_registration_OTP where phone_num=` + phoneNumber + ` and otp = ` + OTP + ` and  status=0  `
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	count := 0
	if err == nil {
		for rows.Next() {

			count++
		}

	} else {
		writeLog("Error in checkOTPused: " + err.Error())

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}
func checkOTPexpird(connection *sql.DB, phoneNumber string, OTP string) bool {
	expiredOTP := GetConfigurationParameter("expiredOTP", "5")

	sqlStatement := `SELECT * FROM user_registration_OTP where phone_num=` + phoneNumber + ` and otp = ` + OTP + ` and  status=0 and TIMESTAMPDIFF(MINUTE,creation_date,CURRENT_TIMESTAMP())<  ` + expiredOTP
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	count := 0
	if err == nil {
		for rows.Next() {

			count++
		}

	} else {
		writeLog("Error in checkOTPexpird: " + err.Error())

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}

func updateStatusOfOTP(connection *sql.DB, phoneNumber string) (err error) {
	sql := `UPDATE user_registration_OTP SET status = -1
	WHERE phone_num =? `

	_, err = connection.Exec(sql, phoneNumber)

	if err != nil {
		writeLog("Error in updateStatusOfOTP: " + err.Error())
	}

	return
}
