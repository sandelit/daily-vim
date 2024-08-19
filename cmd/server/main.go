package main

import (
	"log"
	"net/http"

	"github.com/sandelit/daily-vim/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle main page
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/all-tips", handlers.AllTipsHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
