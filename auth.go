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

func signin(db *sql.DB) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST"{
			type Body struct {
				Username string
				Password string
			}
			var response map[string] string
			response = make(map[string] string)
			body := Body{}
			err := json.NewDecoder(r.Body).Decode(&body)
			defer r.Body.Close()
			user, userId, err := authenticateUser(db, body.Username, body.Password)
			if err != nil {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			} else {
				response["status"] = "OK"
				response["user"] = user
				response["id"] = userId
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func signup(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Body struct {
				Username string
				Password string
				Email string
			}
			var response map[string] string
			response = make(map[string] string)
			body := Body{}
			err := json.NewDecoder(r.Body).Decode(&body)
			defer r.Body.Close()
			user, userId, err := addUser(db, body.Username, body.Password, body.Email)
			if err != nil {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			} else {
				response["status"] = "OK"
				response["user"] = user
				response["id"] = userId
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func addUser(db *sql.DB, username string, password string, email string) (string, string, error) {
	var user string
	query := "select username from USER where username='" + username + "' or email='" + email + "';"
	fmt.Println(query)
	err := db.QueryRow(query).Scan(&user)
	switch {
		case err == sql.ErrNoRows:
			query = "insert into USER set username='" + username + "',email='" + email + "',password='" + password + "';"
			res, err := db.Exec(query)
			if err == nil {
				userId, _ := res.LastInsertId()
				return username, strconv.Itoa(int(userId)), nil
			} else {
				return "", "0", err
			}
		default:
			err = errors.New("User or email already exist")
			return "", "0", err
	}
			
}

func authenticateUser(db *sql.DB, username string, password string) (string, string, error) {
	var user string
	var userId int64
	query := "select username,id from USER where username='" + username + "' and password='" + password + "';"
	fmt.Println(query)
	err := db.QueryRow(query).Scan(&user, &userId)
	if err != nil {
		return "", "0", err
	} else {
		return user, strconv.Itoa(int(userId)), nil
	}
}
