package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var databaseServer, databaseUser, database, password string


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
func InsertReservation(reservation Reservation, reservationTableID string, db *sql.DB) (err error) {

	sql := `INSERT INTO Nowru.reservation (costumer_id,restaurant_id,reservation_table_id,table_type,start_at,end_at)  
	values(?,?,?,?,?,?);
	`
	_, err = db.Exec(sql, reservation.CostumerID, reservation.RestaurantID, reservationTableID, reservation.TableType, reservation.ReservetionStartAt, reservation.ReservationEndAt)

	if err != nil {
		//	writeLog("Error in insertBlackList: " + err.Error())
		fmt.Println(err)

	}
	return err
}

func CheckTableExistence(db *sql.DB, restaurant_id, tableType, start_at string) (tablesIDs []string, existence bool) {

	existence = false
	date, err := time.Parse(time.RFC3339, start_at)
	dateAfterOneHour := date.Add(time.Hour * 1).String()
	dateBeforeOneHour := date.Add(time.Hour * -1).String()
	var sliceOfTableIdOfyourTableType []string

	sqlStatement := `SELECT table_id FROM nowru.tables where restaurant_id = ` + restaurant_id + ` and table_type=` + tableType + `; `
	fmt.Println(sqlStatement)
	rows, err := db.Query(sqlStatement)
	var tableID string
	tableID = "-1"
	if err == nil {
		for rows.Next() {

			rows.Scan(&tableID)
			sliceOfTableIdOfyourTableType = append(sliceOfTableIdOfyourTableType, tableID)

		}
		for index, table := range sliceOfTableIdOfyourTableType {
			sqlStatement1 := `SELECT * FROM nowru.reservation where table_id=` + table + ` and start at between ` + dateBeforeOneHour + `and ` + dateAfterOneHour + `;`
			fmt.Println(sqlStatement)

			rows1, err := db.Query(sqlStatement1)
			count := 0
			if err == nil {
				for rows1.Next() {

					count++
				}

			} else {
				//writeLog("Error in getServices: " + err.Error())

			}
			fmt.Println(err)
			if count > 0 {
				RemoveIndex(sliceOfTableIdOfyourTableType, index)
			}
		}
	} else {
		fmt.Println(err)
		//writeLog("Error in getServices: " + err.Error())
	}
	fmt.Println(err)
	if len(sliceOfTableIdOfyourTableType) > 0 {
		existence = true
	}

	return sliceOfTableIdOfyourTableType, existence

}
