package controller

import (
	"fmt"
	"github.com/gernest/utron"
)

type Login struct {
	*utron.BaseController
	Routes []string
}

func (c *Login) Login() {
	fmt.Println("hello login")
}

func init() {
	utron.RegisterController(&Login{
		Routes: []string{
			"get;/login;Login",
		},
	})
}
