package handlers

import (
    "html/template"
    "net/http"

    "github.com/sandelit/daily-vim/internal/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tip := models.GetRandomTip()

    tmpl, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, tip)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
