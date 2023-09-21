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

func CheckRestaurantExistence(connection *sql.DB, email string) bool {
	sqlStatement := `SELECT * FROM restaurant_registration where email=` + email
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	count := 0
	if err == nil {
		for rows.Next() {

			count++
		}

	} else {
		//writeLog("Error in checkSubcriberExistence: " + err.Error())
		fmt.Println(err)

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}

func insertUserRegistInfo(connection *sql.DB, userRegistrationInfo RestaurantRegistration) (err error) {
	sql := `insert into nowru.restaurant_registration(email,user_password,is_verify,is_activate) values(?,?,?,?);
	`

	_, err = connection.Exec(sql, userRegistrationInfo.Email, userRegistrationInfo.Password, "0", "0")

	if err != nil {
		//writeLog("Error in insertSubscriber: " + err.Error())
		fmt.Println(err)

	}
	return
}

func sendEmail(string) string {
	err := ""
	return err
}
