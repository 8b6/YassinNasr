package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var databaseServer, databaseUser, database, password string

func init() {
	databaseServer = readStringFromConfig("server", "")
	databaseUser = readStringFromConfig("dbuser", "")

	database = readStringFromConfig("database", "")

	password = readStringFromConfig("dbpassword", "")

	println("Inilizalized")

}

func SQLConnection() (result *sql.DB, err error) {

	connectionString := fmt.Sprintf("%v:%v@tcp(%s:2375)/%v?parseTime=true",
		databaseUser, password, databaseServer, database)
	var db *sql.DB
	db, err = sql.Open("mysql", connectionString)
	if err != nil {

		//writeLog("Error in SQLConnection" + err.Error())
		fmt.Print("Error connecting to: ", databaseServer, ": ", err.Error())

	} else {

		fmt.Println("Connected to: ", databaseServer)

	}

	return db, err
}
func getUserIDfromInfo(connection *sql.DB, user_phone_number string) (id string, err error) {

	sqlStatement := `SELECT user_id FROM user_registration where phone_num=` + user_phone_number + `; `
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	user_id := "-1"
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&user_id)

		}

	} else {
		//writeLog("Error in getUserID: " + err.Error())
		fmt.Println(err)
	}
	fmt.Println(err)

	return user_id, err

}

func insertOTP_request(connection *sql.DB, user_phone_number string, user_password, OTP string) (err error) {

	sql := `INSERT INTO nowru.user_registration_OTP (phone_num,password,otp) VALUES (?,?,?);
	`

	_, err = connection.Exec(sql, user_phone_number, user_password, OTP)

	if err != nil {
		fmt.Println(err)
		// writeLog("Error in insertOTP_request: " + err.Error())

	}
	return
}

func getUserID(connection *sql.DB, user_phone_number string) (id string, err error) {

	sqlStatement := `SELECT user_id FROM user_registration where phone_num=` + user_phone_number + `; `
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	user_id := "-1"
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&user_id)

		}

	} else {
		//writeLog("Error in getUserID: " + err.Error())
		fmt.Println(err)
	}
	fmt.Println(err)

	return user_id, err

}

func UpdatePassword(connection *sql.DB, userID, userNewPassword string) error {
	sqlStatement := `update nowru.user_registration set user_password = " ` + userNewPassword + ` " where user_id= ` + userID + `;`
	_, err := connection.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
