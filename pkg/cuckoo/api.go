package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
)

func main() {

    node := []int{1,2,3}
    nodeMap := map[string]interface{}{"node" : node, "hotel": 4  }
    jsonVal, err := json.Marshal(nodeMap)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post("http://127.0.0.1:5000/api", "application/json", bytes.NewBuffer(jsonVal) )
    if err != nil {
        log.Fatal("HTTP Post: ", err)
    }
    defer resp.Body.Close()


    var value map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&value)

    fmt.Printf("%v\n", value)
}
