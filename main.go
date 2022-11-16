package main

import (
	"Ging/framwork"
	"Ging/service/UserService"
	"fmt"
	"net/http"
)

func main() {
	s := framwork.NewHttpServer("FirstServer")

	//s.Route("/", hello)
	//s.Route("/signup", UserService.SignUp)
	//s.RouteWithMethod("GET", "/", hello)
	s.RouteWithMethod(http.MethodPost, "/signup", UserService.SignUp)
	s.RouteWithMethod(http.MethodGet, "/", hello)
	s.Start(":8080")
}

func hello(c *framwork.Context) {
	fmt.Fprintf(c.W, "Hello!")
}
