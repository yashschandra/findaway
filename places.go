package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func addPlace(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Body struct {
				UserId    string
				PlaceName string
			}
			var response map[string]string
			response = make(map[string]string)
			body := Body{}
			err := json.NewDecoder(r.Body).Decode(&body)
			fmt.Println(body)
			defer r.Body.Close()
			place, placeId, err := addNewPlace(db, body.PlaceName, body.UserId)
			if err != nil {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			} else {
				response["status"] = "OK"
				response["place"] = place
				response["id"] = placeId
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func addNewPlace(db *sql.DB, placeName string, userId string) (string, string, error) {
	var place string
	query := "select name from PLACE where name='" + placeName + "';"
	err := db.QueryRow(query).Scan(&place)
	switch {
	case err == sql.ErrNoRows:
		query = "insert into PLACE set name='" + placeName + "',addedByUser=" + userId + ";"
		res, err := db.Exec(query)
		if err == nil {
			placeId, _ := res.LastInsertId()
			return placeName, strconv.Itoa(int(placeId)), nil
		} else {
			return "", "0", err
		}
	default:
		err = errors.New("Place already exist or user does not exist")
		return "", "0", err
	}
}

func getPlaces(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Body struct {
				Ids []string
			}
			type Place struct {
				Id   int64
				Name string
			}
			var places []Place
			var response map[string]interface{}
			response = make(map[string]interface{})
			body := Body{}
			err := json.NewDecoder(r.Body).Decode(&body)
			fmt.Println(body)
			defer r.Body.Close()
			query := "select id,name from PLACE where id IN (" + strings.Join(body.Ids[:], ",") + " ) order by FIELD(id, " + strings.Join(body.Ids[:], ",") + ");"
			fmt.Println(query)
			rows, err := db.Query(query)
			if err == nil {
				for rows.Next() {
					var place Place
					err = rows.Scan(&place.Id, &place.Name)
					if err == nil {
						places = append(places, place)
					}
				}
				response["status"] = "OK"
				response["places"] = places
			} else {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}
