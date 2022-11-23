package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/0xshushu/bitly/handlers"
	"github.com/0xshushu/bitly/database"
)

func main() {
	//initialize database
	handler, err := database.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer handler.DB.Close()
	database.Handler = handler
	defer database.Handler.DB.Close()

	//make a new router
	r := mux.NewRouter().StrictSlash(true)
	
	//set routes
	r.HandleFunc("/", handlers.Index).Methods("GET")
	r.HandleFunc("/shorten", handlers.Shorten).Methods("GET")
	r.HandleFunc("/s/{id}", handlers.Redirect).Methods("GET")

	//listen and serve
	fmt.Println(http.ListenAndServe("127.1:3000", r))
	
}