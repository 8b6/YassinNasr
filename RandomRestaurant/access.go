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

func AllRestaurantIDs(db *sql.DB) (sliceOfrestaurantIDs []string, err error) {

	sqlStatement := `SELECT user_id FROM nowru.restaurant_information ; `
	fmt.Println(sqlStatement)
	rows, err := db.Query(sqlStatement)
	restaurantID := "-1"
	if err == nil {
		for rows.Next() {

			rows.Scan(&restaurantID)
			sliceOfrestaurantIDs = append(sliceOfrestaurantIDs, restaurantID)

		}
		return sliceOfrestaurantIDs, err
	}
	return sliceOfrestaurantIDs, err
}

func Getrandrest(connection *sql.DB, restaurantIDs []string) ([]interface{}, error) {

	sqlStatement := `SELECT category,restaurant_name,profile_pic,pic1,pic2,pic3 FROM restaurant_information where user_id = ` + restaurantIDs[0] + ` or user_id =` + restaurantIDs[1] + `  or user_id = ` + restaurantIDs[2] + `  or user_id = ` + restaurantIDs[3] + `; `
	fmt.Println(sqlStatement)
	rows, err := connection.Query(sqlStatement)
	columnTypes, err := rows.ColumnTypes()
	finalRows := []interface{}{}
	if err != nil {
		return finalRows, err
	}

	count := len(columnTypes)
	//finalRows := []interface{}{}

	for rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {

			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			return finalRows, err
		}

		masterData := map[string]interface{}{}

		for i, v := range columnTypes {

			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}

			masterData[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, masterData)
	}

	//z, err := json.Marshal(finalRows)

	return finalRows, err

}
