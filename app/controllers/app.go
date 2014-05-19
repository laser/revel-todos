package controllers

import "github.com/revel/revel"

type Application struct {
	Gorp
}

func (c Application) Index() revel.Result {
	return c.Redirect(Todos.Index)
}
