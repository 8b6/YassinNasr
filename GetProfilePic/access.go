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

func getProfilePicURL(connection *sql.DB, user_id string) (URL string, err error) {

	sqlStatement := `SELECT profile_pic FROM nowru.user_information where user_id=` + user_id + `; `
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	URL = "-1"
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&URL)

		}

	} else {
		//writeLog("Error in getUserID: " + err.Error())
		fmt.Println(err)
	}
	fmt.Println(err)

	return URL, err

}
