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

func insertUserInfo(connection *sql.DB, insertedUserInfo UserInfo) (err error) {
	sql := `insert into nowru.user_information(user_id,first_name,last_name,gender,birthdate,phone_num) values(?,?,?,?,?,?);
	`

	_, err = connection.Exec(sql, insertedUserInfo.UserID, insertedUserInfo.FirstName, insertedUserInfo.LastName, insertedUserInfo.Gender, insertedUserInfo.Birthdate, insertedUserInfo.PhoneNumber)

	if err != nil {
		//writeLog("Error in insertSubscriber: " + err.Error())
		fmt.Println(err)

	}
	return
}

func insertUserPicURL(connection *sql.DB, userID, url string) bool {
	sqlStatement := `update nowru.user_information set profile_pic = " ` + url + ` " where user_id= ` + userID + `;`
	_, err := connection.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
