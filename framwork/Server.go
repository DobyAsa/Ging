package framwork

import "net/http"

type Server interface {
	Route(pattern string, handlerFunc func(c *Context))
	Start(addr string) error
}

type httpServer struct {
	Name string
}

func (s *httpServer) Route(pattern string, handlerFunc func(c *Context)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(w, r)
		handlerFunc(c)
	})
}

func (s *httpServer) Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}
func NewHttpServer(name string) Server {
	return &httpServer{
		Name: name,
	}
}
