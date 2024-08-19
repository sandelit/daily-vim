package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Content string
		tip     string
	}{
		Title:   "Daily Vim",
		Content: "Welcome to Daily Vim!",
	}

	renderTemplate(w, "web/templates/index.html", data)
}

func AllTipsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Content string
	}{
		Title:   "Daily Vim - All Tips",
		Content: "Welcome to Daily Vim!",
	}

	renderTemplate(w, "web/templates/all-tips.html", data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Content string
	}{
		Title:   "Daily Vim - About",
		Content: "Welcome to Daily Vim!",
	}

	renderTemplate(w, "web/templates/about.html", data)
}

func renderTemplate(w http.ResponseWriter, page string, data interface{}) {
	tmpl, err := template.ParseFiles(page, "web/templates/header.html", "web/templates/github.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
