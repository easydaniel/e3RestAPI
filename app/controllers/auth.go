package controllers

import (
	"github.com/revel/revel"
  "e3plus/app/routes"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	return c.Render()
}

func (c Auth) ValidateLogin() revel.Result {
	return c.Redirect(routes.App.Index())
}
