package controllers

import (
	"github.com/laser/revel-todos/app/models"
	"github.com/laser/revel-todos/app/routes"
	"github.com/revel/revel"
)

type Todos struct {
	Application
}

func (c Todos) Index() revel.Result {
	results, err := c.Txn.Select(models.Todo{}, `select * from todos`)
	if err != nil {
		panic(err)
	}

	completed := 0
	remaining := 0
	var todos []*models.Todo
	for _, r := range results {
		t := r.(*models.Todo)
		todos = append(todos, t)
		if t.Completed {
			completed++
		} else {
			remaining++
		}
	}

	return c.Render(todos, completed, remaining)
}

func (c Todos) ClearCompletedTodos() revel.Result {
	results, err := c.Txn.Select(models.Todo{}, `select * from todos where completed = ?`, true)
	if err != nil {
		panic(err)
	}

	for _, r := range results {
		_, err = c.Txn.Delete(r)
		if err != nil {
			panic(err)
		}
	}

	return c.Redirect(routes.Todos.Index())
}

func (c Todos) DeleteTodo(id int) revel.Result {
	_, err := c.Txn.Delete(&models.Todo{Id: id})
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Todos.Index())
}

func (c Todos) UpdateTodo(id int, title string, completed bool) revel.Result {
	todo := models.Todo{
		Completed: completed,
		Id:        id,
		Title:     title,
	}

	_, err := c.Txn.Update(&todo)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Todos.Index())
}

func (c Todos) CreateTodo(title string) revel.Result {
	todo := models.Todo{
		Completed: false,
		Title:     title,
	}

	err := c.Txn.Insert(&todo)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Todos.Index())
}
