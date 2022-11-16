package main

import (
	"Ging/framwork"
	"Ging/service/UserService"
	"fmt"
	"net/http"
)

func main() {
	s := framwork.NewHttpServer("FirstServer")

	s.Route("/", hello)
	s.Route("/signup", UserService.SignUp)

	s.Start(":8080")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
