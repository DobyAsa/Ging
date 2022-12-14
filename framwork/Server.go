package framwork

import "net/http"

type Server interface {
	Route(pattern string, handlerFunc func(c *Context))
	RouteWithMethod(method string, pattern string,
		handlerFunc func(c *Context))
	Start(addr string) error
}

type httpServer struct {
	Name    string
	Handler Handler
}

func (s *httpServer) RouteWithMethod(method string,
	pattern string, handlerFunc func(c *Context)) {
	s.Handler.Route(method, pattern, handlerFunc)
}

func (s *httpServer) Route(pattern string,
	handlerFunc func(c *Context)) {
	s.RouteWithMethod(http.MethodGet, pattern, handlerFunc)
	s.RouteWithMethod(http.MethodPost, pattern, handlerFunc)
}

func (s *httpServer) Start(addr string) error {
	return http.ListenAndServe(addr, s.Handler)
}
func NewHttpServer(name string) Server {
	return &httpServer{
		Name:    name,
		Handler: NewHandlerBasedOnMap(),
	}
}
