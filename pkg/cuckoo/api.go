package cuckoo

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)
const 
    API = "http://127.0.1:5000/api" 


func GetRoute(node []int, hotelId int) (map[string]interface{}, error) {
    data := map[string]interface{}{"node" : node, "hotel" : hotelId}
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Println("Build Json Error")
        return nil, err
    }

    res, err := http.Post(API, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var route map[string]interface{}

    json.NewDecoder(res.Body).Decode(&route)

    return route, nil
}

