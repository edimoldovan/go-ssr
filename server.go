package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"main/handlers"
	"main/middlewares"
)

//go:embed templates
var embededTemplates embed.FS

//go:embed public
var embededPublic embed.FS

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	// mux/router
	mux := http.NewServeMux()

	// public HTML route middleware stack
	publicHTMLStack := []middlewares.Middleware{
		middlewares.Logger,
	}

	// HTML routes
	mux.HandleFunc("/", middlewares.CompileMiddleware(handlers.Home, publicHTMLStack))

	// static routes, embedded in server binary
	mux.Handle("/public/", handlers.ServeEmbedded(http.FileServer(http.FS(embededPublic))))

	// HTTP server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
