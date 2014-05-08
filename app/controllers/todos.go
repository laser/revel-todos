package controllers

import "github.com/revel/revel"

type Todos struct {
	Application
}

func (c Todos) Index() revel.Result {
	todos := []string{"Call dad"}
	return c.Render(todos)
}
