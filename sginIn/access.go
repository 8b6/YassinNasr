package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var databaseServer, databaseUser, database, password string

func init() {
	databaseServer = GetConfigurationParameter("server", "")
	databaseUser = GetConfigurationParameter("dbuser", "")

	database = GetConfigurationParameter("database", "")

	password = GetConfigurationParameter("dbpassword", "")
	db, _ = SQLConnection()

	println("Inilizalized")

}
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

func insertUserInformation(connection *sql.DB, msisdn string, servcieCode string) (err error) {
	sql := `insert into nowru.user_information(user_first_name,user_last_name,username,gender,birthdate,food_preference,user_profile_pic) values(?,?,?,?,?,?,?,?);
	`

	_, err = connection.Exec(sql, msisdn, servcieCode)

	if err != nil {
		writeLog("Error in insertSubscriber: " + err.Error())
		fmt.Println(err)

	}
	return
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
		writeLog("Error in checkSubcriberExistence: " + err.Error())

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}
