package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Way struct {
	FromPlaceId int64
	ToPlaceId   int64
	FromPlace   string
	ToPlace     string
	Mode        string
	Distance    int
	Cost        int
	WayId       int64
}

func way(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			type Body struct {
				From int64
				To   int64
			}
			var response map[string]interface{}
			response = make(map[string]interface{})
			body := Body{}
			_ = json.NewDecoder(r.Body).Decode(&body)
			defer r.Body.Close()
			path, err := findAWay(db, body.From, body.To)
			if err == nil {
				response["status"] = "OK"
				response["path"] = path
			} else {
				response["status"] = "Not OK"
				response["error"] = err.Error()
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func dummyWay(from int64) Way {
	var way Way
	way.Cost = 0
	way.Distance = 0
	way.FromPlace = ""
	way.FromPlaceId = 0
	way.Mode = ""
	way.ToPlace = ""
	way.ToPlaceId = from
	way.WayId = 0
	return way
}

func findAWay(db *sql.DB, from int64, to int64) ([]Way, error) {
	var ways []Way
	var path []Way
	query := "select fromPlace.name,fromPlace.id,toPlace.name,toPlace.id,way.cost,mode.name,route.distance,way.id from WAY way join ROUTE route on way.route=route.id join PLACE fromPlace on fromPlace.id=route.fromPlace join PLACE toPlace on toPlace.id=route.toPlace join MODE mode on mode.id=way.mode;"
	rows, err := db.Query(query)
	if err == nil {
		for rows.Next() {
			var way Way
			err = rows.Scan(&way.FromPlace, &way.FromPlaceId, &way.ToPlace, &way.ToPlaceId, &way.Cost, &way.Mode, &way.Distance, &way.WayId)
			if err == nil {
				ways = append(ways, way)
			}
		}
		var visited []int64
		parent := make([]Way, len(ways)+1)
		var queue []Way
		var ele Way
		dWay := dummyWay(from)
		queue = append(queue, dWay)
		for len(queue) > 0 {
			ele = queue[0]
			queue = queue[1:]
			if !isPresent(ele.WayId, visited) {
				visited = append(visited, ele.WayId)
				for _, w := range ways {
					if w.FromPlaceId == ele.ToPlaceId {
						if !isPresent(w.WayId, visited) {
							parent[w.WayId] = ele
							queue = append(queue, w)
							if w.ToPlaceId == to {
								path = trackPath(parent, w, ways)
								return path, nil
							}
						}
					}
				}
			}
		}
		return nil, nil
	} else {
		return path, err
	}
}

func isPresent(ele int64, visited []int64) bool {
	for _, e := range visited {
		if ele == e {
			return true
		}
	}
	return false
}

func reverseWay(arr []Way) []Way {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func trackPath(parent []Way, way Way, ways []Way) []Way {
	var path []Way
	path = append(path, way)
	to := way.WayId
	for parent[to].WayId != 0 {
		path = append(path, parent[to])
		to = parent[to].WayId
	}
	path = reverseWay(path)
	return path
}
