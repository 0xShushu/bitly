package handlers

import (
	"net/http"
	"html/template"
)

//handle all 404 status codes
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	//get template file
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//execute template
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}