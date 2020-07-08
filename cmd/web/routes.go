package main
import (
    "net/http"
)

func (app *application) routes() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/", app.home)

    fileserver := http.FileServer(http.Dir("ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileserver) )
    mux.HandleFunc("/route", app.route)

    return mux
}
