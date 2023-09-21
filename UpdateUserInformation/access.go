package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

func updateTableUserRegist(connection *sql.DB, userInfo UserInfo) bool {
	sqlStatement := `update nowru.user_registration set 
	phone_num = ?,
	user_password = ? 
	where user_id = ?;`
	_, err := connection.Exec(sqlStatement, userInfo.PhoneNumber, userInfo.Password, userInfo.UserID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func updateTableUserinfo(connection *sql.DB, userInfo UserInfo) bool {
	sqlStatement := `update nowru.user_information set 
	first_name = ?,
	last_name = ?,
	gender = ?,
	birthdate = ?,
	phone_num = ? 
	where user_id = ?;`
	_, err := connection.Exec(sqlStatement, userInfo.FirstName, userInfo.LastName, userInfo.Gender, userInfo.Birthdate, userInfo.PhoneNumber, userInfo.UserID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
