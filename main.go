package main

import (
	"github.com/christoperBar/Todolist_api/controllers/todocontroller"
	"github.com/christoperBar/Todolist_api/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	todo := api.Group("/todoLists")

	todo.Get("/", todocontroller.AllTodoLists)
	todo.Get("/:id", todocontroller.GetTodoList)
	todo.Post("/", todocontroller.AddTodoList)
	todo.Put("/:id", todocontroller.UpdateTodoList)
	todo.Delete("/:id", todocontroller.DeleteTodoList)

	app.Listen(":8000")
}
