package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var Tmpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	if err := Tmpl.ExecuteTemplate(w, "home", map[string]interface{}{
		"Title": "Web SSR with Go",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
