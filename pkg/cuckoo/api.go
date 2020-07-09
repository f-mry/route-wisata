package cuckoo

import (
	"bytes"
	"encoding/json"
	// "fmt"
	// "io/ioutil"
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

// func main() {

//     node := []int{1,2,3}
//     nodeMap := map[string]interface{}{"node" : node, "hotel": 4  }
//     jsonVal, err := json.Marshal(nodeMap)
//     if err != nil {
//         log.Fatal(err)
//     }

//     resp, err := http.Post("http://127.0.0.1:5000/api", "application/json", bytes.NewBuffer(jsonVal) )
//     if err != nil {
//         log.Fatal("HTTP Post: ", err)
//     }
//     defer resp.Body.Close()


//     var value map[string]interface{}
//     json.NewDecoder(resp.Body).Decode(&value)

//     fmt.Printf("%v\n", value)
// }
