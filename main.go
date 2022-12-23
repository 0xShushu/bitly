package main

import (
	"net/http"
	"fmt"
	"github.com/0xshushu/bitly/database"
	"github.com/0xshushu/bitly/server"
)

//cringe

func main() {
	//initialize database
	handler, err := database.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	database.Handler = handler
	defer database.Handler.DB.Close()

	//make a new router
	s := server.NewServer()
	
	//mount handlers
	s.MountHandlers()

	//listen and serve
	fmt.Println(http.ListenAndServe(":3000", s.Router))	
}
