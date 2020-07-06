package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from server")
}


func main() {
    
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    log.Println("Serving on :4040")
    err := http.ListenAndServe(":4040", mux)
    if err != nil {
        log.Fatal(err)
    }
}
