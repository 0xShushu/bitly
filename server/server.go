package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/0xshushu/bitly/handlers"
)

//server struct
type Server struct {
	Router *chi.Mux
}

//function to initialize a new server
func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

//function to mount handlers
func (s *Server) MountHandlers() {
	//set middleware
	s.Router.Use(middleware.Logger)

	//set 404 route
	s.Router.NotFound(handlers.NotFoundHandler)

	//set / route
	s.Router.Get("/", handlers.Index)
	
	//set /shorten route
	s.Router.Get("/shorten", handlers.Shorten)
	s.Router.Get("/s/{id}", handlers.Redirect)
}