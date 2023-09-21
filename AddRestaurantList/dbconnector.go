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

		//writeLog("Error in SQLConnection" + err.Error())
		fmt.Print("Error connecting to: ", databaseServer, ": ", err.Error())

	} else {

		fmt.Println("Connected to: ", databaseServer)

	}

	return db, err

}
func CloseConnection(db sql.DB) {
	db.Close()
}

func checkUserRestaurantLimit(userId string, connection *sql.DB) bool {
	sqlStatement := `SELECT userlimit FROM userrestaurantlist  where userid=` + userId
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	count := 0
	if err == nil {
		for rows.Next() {

			count++
		}

	} else {
		//writeLog("Error in getServices: " + err.Error())

	}
	fmt.Println(err)
	if count > 0 {
		return true
	}

	return false
}

func insertedFavRestaurant(insertedFavRestaurantId, userId, db *sql.DB) (err error) {

	sql := `INSERT INTO Nowru.favRestaurant (userId,restaurantId)  
	values(?,?);
	`
	//change the type of execute the sql statement
	_, err = db.Exec(sql, insertedFavRestaurantId, userId)

	if err != nil {
		//	writeLog("Error in insertBlackList: " + err.Error())
		fmt.Println(err)

	}
	return
}
func updateService(db sql.DB, insertedService UpdateService, serivceId string) error {

	ServiceAfterUpdate := ServiceAfterUpdate{"0", "0", "0", "0", "0"}
	if insertedService.ServicePrice != "" {
		sqlStatement := `update services set service_price =` + insertedService.ServicePrice + ` where services_id= ` + serivceId + `;`
		_, err := db.Exec(sqlStatement)
		// err := UpdateServiceColoum("service_price", insertedService.ServicePrice, serivceId, db)

		if err != nil {
			fmt.Println(err)
		} else {
			ServiceAfterUpdate.ServicePrice = "updated"
			fmt.Println(ServiceAfterUpdate)

		}
	}
	if insertedService.Description != "" {
		sqlStatement := `UPDATE tpay.services
		SET
		description ="` + insertedService.Description + `" 
		where services_id= ` + serivceId + `;`
		fmt.Println(sqlStatement)
		_, err := db.Exec(sqlStatement)
		// err := UpdateServiceColoum("description", insertedService.Description, serivceId, db)

		if err != nil {
			fmt.Println(err)
		} else {
			ServiceAfterUpdate.Description = "updated"
			fmt.Println(ServiceAfterUpdate)
		}
	}
	if insertedService.HomePage != "" {
		// sqlStatement := `update services set home_page=` + insertedService.HomePage + ` where services_id= ` + serivceId + `;`
		sqlStatement := `UPDATE tpay.services
		SET
		home_page ="` + insertedService.HomePage + `" 
		where services_id= ` + serivceId + `;`
		fmt.Println(sqlStatement)
		_, err := db.Exec(sqlStatement)
		// err := UpdateServiceColoum("home_page", insertedService.HomePage, serivceId, db)

		if err != nil {
			fmt.Println(err)

		} else {
			ServiceAfterUpdate.HomePage = "updated"
			fmt.Println(ServiceAfterUpdate)
		}
	}
	if insertedService.ServiceType != "" {
		err := UpdateServiceColoum("service_type", insertedService.ServiceType, serivceId, db)

		if err != nil {
			fmt.Println(err)
		} else {
			ServiceAfterUpdate.ServiceType = "updated"
			fmt.Println(ServiceAfterUpdate)
		}
	}
	return nil
}

func UpdateServiceColoum(updatedColoumName, updatedValue, services_id string, db sql.DB) (err error) {

	sqlStatement := `update services set ` + updatedColoumName + ` =` + updatedValue + ` where services_id= ` + services_id
	_, err = db.Exec(sqlStatement)

	return
}

func getServiceID(connection *sql.DB, service_code string) (id string, err error) {

	sqlStatement := `SELECT services_id FROM services where service_code=` + service_code + `; `
	fmt.Println(sqlStatement)

	rows, err := connection.Query(sqlStatement)
	var services_id string
	services_id = "-1"
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&services_id)

		}

	} else {
		writeLog("Error in getServiceID: " + err.Error())
	}
	fmt.Println(err)

	return services_id, err

}
