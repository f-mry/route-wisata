package main

import (
	"fmt"
	"krisna/pkg/models"
	"net/http"
	"strconv"
	"text/template"
    // "net/url"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.infoLog.Printf("URL Not Found: %s", r.URL.Path)
        fmt.Fprintf(w, "Page Not Found: %s", r.URL.RawPath)
        return
    }

    page := r.URL.Query().Get("page")

    if page == "" {
        page = "0"
    }

    index, err := strconv.Atoi(page)
    if err != nil {
        app.infoLog.Printf("Query error: %v", err)
        fmt.Fprintf(w, "Page Not Found: %s", r.URL.RawPath)
        return
    }

    wisata, err := app.wisata.Get(index)
    if err != nil {
        app.infoLog.Println(err)
        app.serverError(w, err)
        return
    }
    

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

func (app *application) route(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()["node"]
    for _, q := range(query) {
        fmt.Println(q)
    }

    fmt.Fprintf(w, "Query: %v", query)
}

