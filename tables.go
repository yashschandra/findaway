package main

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func executeDBQuery(db *sql.DB, query string) error {
	fmt.Println(query)
	stmt, err := db.Prepare(query)
	_, err = stmt.Exec()
	checkError(err)
	return err
}

func createTable(db *sql.DB, tableName string, fields string) {
	query := "create table " + tableName + " (" + fields + ");"
	err := executeDBQuery(db, query)
	checkError(err)
}

func dropTable(db *sql.DB, tableName string){
	query := "drop table " + tableName + ";"
	err := executeDBQuery(db, query)
	checkError(err)
}

func createUserTable(db *sql.DB) {
	tableName := "USER"
	fields := "id int AUTO_INCREMENT PRIMARY KEY, email varchar(30) NOT NULL UNIQUE, username varchar(30) NOT NULL UNIQUE, password varchar(30) NOT NULL"
	createTable(db, tableName, fields)
}

func createPlaceTable(db *sql.DB){
	tableName := "PLACE"
	fields := "id int AUTO_INCREMENT PRIMARY KEY, name varchar(30) NOT NULL, addedByUser int NOT NULL, FOREIGN KEY (addedByUser) REFERENCES USER(id)"
	createTable(db, tableName, fields)
}

func createModeTable(db *sql.DB){
	tableName := "MODE"
	fields := "id int AUTO_INCREMENT PRIMARY KEY, name varchar(30) NOT NULL, addedByUser int NOT NULL, FOREIGN KEY (addedByUser) REFERENCES USER(id)"
	createTable(db, tableName, fields)
}

func createRouteTable(db *sql.DB){
	tableName := "ROUTE"
	fields := "id int AUTO_INCREMENT PRIMARY KEY, addedbyUser int NOT NULL, fromPlace int NOT NULL, toPlace int NOT NULL, distance int NOT NULL, FOREIGN KEY (addedByUser) REFERENCES USER(id), FOREIGN KEY (fromPlace) REFERENCES PLACE(id), FOREIGN KEY (toPlace) REFERENCES PLACE(id)"
	createTable(db, tableName, fields)
}

func createWayTable(db *sql.DB){
	tableName := "WAY"
	fields := "id int AUTO_INCREMENT PRIMARY KEY, route int NOT NULL, addedByUser int NOT NULL, cost int NOT NULL, mode int NOT NULL, FOREIGN KEY (route) REFERENCES ROUTE(id), FOREIGN KEY (addedByUser) REFERENCES USER(id), FOREIGN KEY (mode) REFERENCES MODE(id)"
	createTable(db, tableName, fields)
}

func createTables(db *sql.DB) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		createUserTable(db)
		createPlaceTable(db)
		createModeTable(db)
		createRouteTable(db)
		createWayTable(db)
	}
}

func dropTables(db *sql.DB) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		tableNames := [...]string{"WAY", "MODE", "ROUTE", "PLACE", "USER"}
		for i := 0; i<len(tableNames); i++ {
			dropTable(db, tableNames[i])
		}
	}
}

func checkError(err error) {
	if err != nil{
		fmt.Println(err.Error())
	}
}
