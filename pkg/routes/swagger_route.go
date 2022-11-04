package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(a *fiber.App) {
	router := a.Group("/swagger")

	router.Get("*", swagger.HandlerDefault)
}
