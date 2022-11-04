package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opchav/todos-go-api/app/controllers"
	"github.com/opchav/todos-go-api/pkg/middleware"
)

func PrivateRoutes(a *fiber.App) {
  router := a.Group("api/v1")

  router.Post("/todos", middleware.JWTProtected(), controllers.CreateTodo)
  router.Put("/todos", middleware.JWTProtected(), controllers.UpdateTodo)
  router.Delete("/todos", middleware.JWTProtected(), controllers.DeleteTodo)
}
