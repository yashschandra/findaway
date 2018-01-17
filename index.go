package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func indexPlaces(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Place struct {
				Id   int64
				Name string
			}
			query := "select id,name from PLACE;"
			rows, err := db.Query(query)
			if err == nil {
				for rows.Next() {
					var place Place
					err = rows.Scan(&place.Id, &place.Name)
					if err == nil {
						data, err := json.Marshal(place)
						if err == nil {
							url := "http://localhost:9200/findaway/places/" + strconv.Itoa(int(place.Id))
							response, _ := http.Post(url, "application/json", bytes.NewBuffer(data))
							fmt.Println(response.Body)
						}
					}
				}
			}
		}
	}
}

func searchPlaces(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Body struct {
				Search string
			}
			body := Body{}
			err := json.NewDecoder(r.Body).Decode(&body)
			if err == nil {
				fmt.Println(body)
				defer r.Body.Close()
				url := "http://localhost:9200/findaway/places/_search?q=Name:*" + body.Search + "*"
				res, _ := http.Get(url)
				defer res.Body.Close()
				var response map[string]interface{}
				data, err := ioutil.ReadAll(res.Body)
				if err == nil {
					jsonData := string(data)
					json.Unmarshal([]byte(jsonData), &response)
					response["status"] = "OK"
				} else {
					response["status"] = "Not OK"
				}
				json.NewEncoder(w).Encode(response)
			} else {
				fmt.Println("error : " + err.Error())
			}
		}
	}
}
