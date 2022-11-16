package framwork

import (
	"fmt"
	"net/http"
)

type Handler interface {
	Route(method string, pattern string, handlerFunc func(c *Context))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HandlerBasedOnMap struct {
	Roads map[string]func(c *Context)
}

func (h *HandlerBasedOnMap) Route(method string,
	pattern string, handlerFunc func(c *Context)) {
	key := h.key(method, pattern)
	h.Roads[key] = handlerFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.Roads[key]; ok {
		c := NewContext(w, r)
		handler(c)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Have no match path")
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return fmt.Sprintf("%s#%s", method, pattern)
}

func NewHandlerBasedOnMap() *HandlerBasedOnMap {
	return &HandlerBasedOnMap{
		Roads: make(map[string]func(c *Context)),
	}
}
