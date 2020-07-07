package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

    "krisna/pkg/mysql"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from server")
}

type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
    wisata *mysql.WisataModel
}


func main() {
    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
    infoLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

    db, err := mysql.OpenDBConnection()
    if err != nil {
        errorLog.Fatal(err)
    }
    defer db.Close()

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        wisata: &mysql.WisataModel{DB: db},
    }

    server := &http.Server{
        Addr: ":4040",
        ErrorLog: app.errorLog,
        Handler: app.routes(),
    }

    err = server.ListenAndServe()
    if err != nil {
        errorLog.Fatal(err)
    }

}
