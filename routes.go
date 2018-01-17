package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func addRoute(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Body struct {
				FromPlace string
				ToPlace   string
				UserId    string
				Distance  string
				Cost      string
				Mode      string
			}
			var response map[string]string
			response = make(map[string]string)
			body := Body{}
			err := json.NewDecoder(r.Body).Decode(&body)
			fmt.Println(body)
			defer r.Body.Close()
			routeId, err := addNewRoute(db, body.FromPlace, body.ToPlace, body.Distance, body.UserId, body.Cost, body.Mode)
			if err != nil {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			} else {
				response["status"] = "OK"
				response["id"] = routeId
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func addNewRoute(db *sql.DB, fromPlace string, toPlace string, distance string, userId string, cost string, mode string) (string, error) {
	var routeId int64
	var wayId string
	query := "select id from ROUTE where fromPlace=" + fromPlace + " and toPlace=" + toPlace + ";"
	fmt.Println(query)
	err := db.QueryRow(query).Scan(&routeId)
	fmt.Println(routeId)
	switch {
	case err == sql.ErrNoRows:
		query = "insert into ROUTE set fromPlace=" + fromPlace + ",toPlace=" + toPlace + ",addedByUser=" + userId + ",distance=" + distance + ";"
		res, err := db.Exec(query)
		if err == nil {
			routeId, _ = res.LastInsertId()
		} else {
			return "0", err
		}
	default:
		fmt.Println("Route already exist id - " + strconv.Itoa(int(routeId)))
	}
	wayId, err = addWay(db, strconv.Itoa(int(routeId)), mode, cost, distance, userId)
	return wayId, err
}

func addWay(db *sql.DB, routeId string, mode string, cost string, distance string, userId string) (string, error) {
	query := "insert into WAY set route=" + routeId + ",cost=" + cost + ",mode=" + mode + ",addedByUser=" + userId + ";"
	res, err := db.Exec(query)
	if err == nil {
		wayId, _ := res.LastInsertId()
		return strconv.Itoa(int(wayId)), nil
	} else {
		fmt.Println("return error:" + err.Error())
		return "0", err
	}
}
