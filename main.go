package main

import "net/http"

func main() {
	server := NewServer(":3000")

	server.Handle(http.MethodGet, "/", HandleRoot)
	server.Handle(http.MethodPost, "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle(http.MethodPost, "/create", server.AddMiddleware(PostHandle, Logging()))
	server.Handle(http.MethodPost, "/user", server.AddMiddleware(UserPostHandle, Logging()))

	server.Listen()
}
