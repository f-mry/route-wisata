package main

import (
	"fmt"
	"krisna/pkg/models"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.infoLog.Printf("URL Not Found: %s", r.URL.Path)
        fmt.Fprintf(w, "Page Not Found: %s", r.URL.RawPath)
        return
    }
    // fmt.Fprintf(w, "Hello Aha")
    wisata, err := app.wisata.Get()
    if err != nil {
        app.infoLog.Println(err)
        app.serverError(w, err)
        return
    }
    fmt.Println(wisata)

    template_files := []string{
        "ui/html/home.page.html",
    }

    ts, err := template.ParseFiles(template_files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    // type Data struct {
    //     listWisata []*models.InfoWisata
    // }

    var Data struct{
        ListWisata []*models.InfoWisata
    }

    Data.ListWisata = wisata

    err = ts.Execute(w, Data )
    if err != nil {
        app.serverError(w, err)
        return
    }
}
