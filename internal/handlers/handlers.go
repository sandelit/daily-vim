package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/sandelit/daily-vim/internal/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Tip   models.Tip
	}{
		Title: "Daily Vim",
		Tip:   models.GetRandomTip(),
	}

	renderTemplate(w, "web/templates/index.html", data)
}

func AllTipsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Tips  []models.Tip
	}{
		Title: "Daily Vim - All Tips",
		Tips:  models.GetAllTips(),
	}

	renderTemplate(w, "web/templates/all-tips.html", data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Daily Vim - About",
	}

	renderTemplate(w, "web/templates/about.html", data)
}

func renderTemplate(w http.ResponseWriter, page string, data interface{}) {
	funcMap := template.FuncMap{
		"processContent": processContent,
	}

	// Get the base name of the template file
	templateName := filepath.Base(page)

	// Parse the main template file
	tmpl, err := template.New(templateName).Funcs(funcMap).ParseFiles(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse and associate additional template files
	tmpl, err = tmpl.ParseFiles("web/templates/header.html", "web/templates/github.html", "web/templates/tip.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the main template
	err = tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func processContent(content string) template.HTML {
	re := regexp.MustCompile(`---(.*?)---`)
	processed := re.ReplaceAllString(content, `<span class="vim-command">$1</span>`)

	return template.HTML(processed)
}
