package main

import (
	"fmt"
	"krisna/pkg/cuckoo"
	"krisna/pkg/models"
	"net/http"
	"strconv"
	"text/template"
    // "html/template"

	"github.com/mitchellh/mapstructure"
	// "net/url"
    "strings"
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

    hotel, err := app.hotel.Get()
    if err != nil {
        app.infoLog.Println(err)
    }

    FuncMap := template.FuncMap{
        "trim" : func(str string) string{
            lastSpace := strings.LastIndexByte(str[0:200], ' ')
            return str[0:lastSpace + 1]
        },
    }
    

    template_files := []string{
        "ui/html/home.page.html",
    }

    ts, err := template.New("home.page.html").Funcs(FuncMap).ParseFiles(template_files...)

    // ts, err := template.ParseFiles(template_files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    // type Data struct {
    //     listWisata []*models.InfoWisata
    // }

    var Data struct{
        ListWisata []*models.InfoWisata
        ListHotel []*models.InfoHotel
    }

    Data.ListHotel = hotel
    Data.ListWisata = wisata


    err = ts.Execute(w, Data)
    if err != nil {
        app.serverError(w, err)
        return
    }
}

func (app *application) route(w http.ResponseWriter, r *http.Request) {
    nodeString := r.URL.Query()["node"]
    hotelString := r.URL.Query().Get("hotel")
    hotel, _ := strconv.Atoi(hotelString)
    node := []int{}

    for _, q := range(nodeString) {
        // fmt.Printf("%v, %T\n", q,q)
        nodeId, err := strconv.Atoi(q)
        if err != nil {
            app.infoLog.Printf("Node parse error: %v", err)
        }
        node = append(node, nodeId)
    }

    route, err := cuckoo.GetRoute(node, hotel)
    if err != nil {
        app.infoLog.Println(err)
        return
    }

    var Result models.Route
    err = mapstructure.Decode(route, &Result)
    if err != nil {
        app.infoLog.Println(err)
        return
    }

    // fmt.Printf("%v\n", Result)
    // fmt.Fprintf(w, "Query: %v\nHotel: %v\nData: \n%#v", node, hotel, Result)

    template_files := []string{
        "ui/html/route.page.html",
    }

    FuncMap := template.FuncMap{
        "inc" : func(i int) int {
            return i + 1
        },
        "wypts" : func(i int) rune {
            return rune('A' - 1 + i)
        },
    }
    
    ts, err := template.New("route.page.html").Funcs(FuncMap).ParseFiles(template_files...)
    // ts, err := template.ParseFiles(template_files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.Execute(w, Result)
    if err != nil {
        app.serverError(w, err)
        return
    }
}

