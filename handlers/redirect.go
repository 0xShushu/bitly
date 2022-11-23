package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/0xshushu/bitly/database"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	//get id from url variables
	vars := mux.Vars(r)
	id := vars["id"]

	//get url from sqlite db
	url, err := database.Select(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//redirect to the url
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}