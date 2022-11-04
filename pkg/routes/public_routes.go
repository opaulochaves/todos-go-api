package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opchav/todos-go-api/app/controllers"
)

func PublicRoutes(a *fiber.App) {
  router := a.Group("/api/v1")

  router.Get("/todos", controllers.GetTodos)
  router.Get("/todos/:id", controllers.GetTodo)
  router.Get("/token/new", controllers.GetNewAccessToken)
}
