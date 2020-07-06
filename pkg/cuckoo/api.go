package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
    query := `{"idWisata":[{"id":"0"},{"id":"1"},{"id":"2"},{"id":"3"}],"idhotel":["1"],"degree":[1,1,1]}`

    resp, err := http.Get("http://127.0.0.1:5000/cuckoo/" + query)
    if err != nil {
        fmt.Println(err)
    }

    // fmt.Print(resp.Body)
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    var value map[string]interface{}
    err = json.Unmarshal(data, &value)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(value)



}
