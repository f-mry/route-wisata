package main

import (
	"fmt"
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

    template_files := []string{
        "ui/html/home.page.html",
    }

    ts, err := template.ParseFiles(template_files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        app.serverError(w, err)
        return
    }
}
