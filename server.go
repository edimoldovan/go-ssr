package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"main/handlers"
	"main/middlewares"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

//go:embed templates
var embededTemplates embed.FS

//go:embed public
var embededPublic embed.FS

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	// router
	router := httprouter.New()

	// middlewares
	chain := alice.New(middlewares.Logger)

	// HTML routes
	router.GET("/", middlewares.Wrapper(chain.ThenFunc(handlers.Home)))

	// static routes, embedded in server binary
	if public, err := fs.Sub(embededPublic, "public"); err == nil {
		router.Handler("GET", "/public/*filepath", http.StripPrefix("/public", http.FileServer(http.FS(public))))
	} else {
		panic(err)
	}

	// HTTP server
	log.Fatal(http.ListenAndServe(":8000", router))
}
