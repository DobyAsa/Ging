package framwork

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

func (c *Context) ReadJSON(data interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		fmt.Fprintln(c.W, err)
		return err
	}

	return json.Unmarshal(body, data)
}

func (c *Context) WriteJSON(status int, data interface{}) error {
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(c.W, err)
	}

	c.W.WriteHeader(status)
	_, err = c.W.Write(res)
	return err
}

func (c *Context) OkJSON(data interface{}) error {
	return c.WriteJSON(http.StatusOK, data)
}
