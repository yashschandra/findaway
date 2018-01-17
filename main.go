package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	}
}

func main() {
	db, err := sql.Open("mysql", "root:rahul@tcp(127.0.0.1:3306)/findaway")
	if err == nil {
		defer db.Close()
		http.HandleFunc("/", index())
		http.HandleFunc("/auth/signin", signin(db))
		http.HandleFunc("/auth/signup", signup(db))
		http.HandleFunc("/places/add", addPlace(db))
		http.HandleFunc("/places/get", getPlaces(db))
		http.HandleFunc("/routes/add", addRoute(db))
		http.HandleFunc("/modes/add", addMode(db))
		http.HandleFunc("/modes/get", getModes(db))
		http.HandleFunc("/createtables", createTables(db))
		http.HandleFunc("/droptables", dropTables(db))
		http.HandleFunc("/indexplaces", indexPlaces(db))
		http.HandleFunc("/searchplaces", searchPlaces(db))
		http.HandleFunc("/way", way(db))
		log.Fatal(http.ListenAndServe(":9001", nil))
	} else {
		fmt.Println("error starting db : " + err.Error())
	}
}
