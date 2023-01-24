package main

import (
	"html/template"
	"log"
	"main/handlers"
	"main/middlewares"
	"main/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func main() {
	// read template files
	templateFiles := utilities.GetTemplates()
	// parse template files
	handlers.Tmpl, _ = template.ParseFiles(templateFiles...)
	// router
	router := httprouter.New()

	// middlewares
	chain := alice.New(middlewares.Logger)

	// HTML routes
	router.GET("/", middlewares.Wrapper(chain.ThenFunc(handlers.Home)))

	// static routes
	f := utilities.GetExecutablePath()
	router.ServeFiles("/public/*filepath", http.Dir(f+"public"))

	// HTTP server
	log.Fatal(http.ListenAndServe(":8000", router))
}
