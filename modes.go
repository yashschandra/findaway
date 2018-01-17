package main

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

func addMode(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Body struct {
				UserId string
				ModeName string 
			}
			var response map[string] string
			response = make(map[string] string)
			body := Body{}
			err := json.NewDecoder(r.Body).Decode(&body)
			fmt.Println(body)
			defer r.Body.Close()
			modeName, modeId, err := addNewMode(db, body.ModeName, body.UserId)
			if err != nil {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			} else {
				response["status"] = "OK"
				response["mode"] = modeName
				response["id"] = modeId
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func addNewMode(db *sql.DB, modeName string, userId string) (string, string, error) {
	var mode string
	query := "select name from MODE where name='" + modeName + "';"
	err := db.QueryRow(query).Scan(&mode)
	switch {
		case err == sql.ErrNoRows:
			query = "insert into MODE set name='" + modeName + "',addedByUser=" + userId + ";"
			res, err := db.Exec(query)
			if err == nil {
				modeId, _ := res.LastInsertId()
				return modeName, strconv.Itoa(int(modeId)), nil
			} else {
				return "", "0", err
			}
		default:
			err = errors.New("Place already exist or user does not exist")
			return "", "0", err
	}
}

func getModes(db *sql.DB) func (w http.ResponseWriter, r *http.Request){
	return func (w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
			type Mode struct {
				Id int64
				Name string
			}
			var modes []Mode
			var response map[string] interface{}
			response = make(map[string] interface{})
			query := "select id,name from MODE;"
			rows, err := db.Query(query)
			if err == nil {
				for rows.Next() {
					var mode Mode
					err = rows.Scan(&mode.Id, &mode.Name)
					if err == nil {
						modes = append(modes, mode)
					}
				}
				response["status"] = "OK"
				response["modes"] = modes
			} else {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}
