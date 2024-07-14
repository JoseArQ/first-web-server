package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {

	_, exist := s.router.rules[path]

	if !exist {
		fmt.Println("creating new rule...")
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(hf http.HandlerFunc, middlawares ...Middleware) http.HandlerFunc {
	for _, middlaware := range middlawares {
		hf = middlaware(hf)
	}

	return hf
}

func (s *Server) Listen() error {

	err := http.ListenAndServe(s.port, s.router)
	if err != nil {
		return err
	}
	log.Println("Server listen on port ", s.port)
	return nil
}
