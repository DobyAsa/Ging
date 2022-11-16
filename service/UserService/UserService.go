package UserService

import (
	"Ging/framwork"
	"fmt"
)

type SignUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

func SignUp(c *framwork.Context) {
	req := &SignUpReq{}

	c.ReadJSON(req)
	fmt.Fprintln(c.W, "Sign up success!")
	fmt.Fprintf(c.W, "%v", req)
}
