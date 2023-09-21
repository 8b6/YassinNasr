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

func getRestaurantID(connection *sql.DB, rest string) (id string, err error) {

	sqlStatement := `SELECT user_id FROM nowru.restaurant_information WHERE restaurant_name =?`
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement, rest)
	restaurant_id := "-1"
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&restaurant_id)

		}

	} else {
		//writeLog("Error in getUserID: " + err.Error())
		fmt.Println(err)
	}
	fmt.Println(err)

	return restaurant_id, err

}

func insertUserReview(connection *sql.DB, userID, restaurantID, userReview, reviewPic string) bool {
	sql := `insert into nowru.review(user_id,restaurant_id,review,review_pic) values(?,?,?,?);
	`

	_, err := connection.Exec(sql, userID, restaurantID, userReview, reviewPic)

	if err != nil {
		//writeLog("Error in insertSubscriber: " + err.Error())
		fmt.Println(err)
		return false

	}
	return true
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
