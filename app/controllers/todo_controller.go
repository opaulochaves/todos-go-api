package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/opchav/todos-go-api/app/models"
	"github.com/opchav/todos-go-api/pkg/utils"
	"github.com/opchav/todos-go-api/platform/database"
)

// GetTodos func gets all existing todos.
// @Description Get all existing todos.
// @Summary get all existing todos.
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {array} models.Todo
// @Router /api/v1/todos [get]
func GetTodos(c *fiber.Ctx) error {
	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	todos, err := db.GetTodos()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "todos were not found",
			"count": 0,
			"todos": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(todos),
		"todos": todos,
	})
}

// GetTodo func gets todo by given ID or 404 error.
// @Description Get book by given ID.
// @Summary get book by given ID.
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /api/v1/todos/{id} [get]
func GetTodo(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	todo, err := db.GetTodo(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "todo with the given ID is not found",
			"todo":  nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"todo":  todo,
	})
}

// CreateTodo func for creating a new todo.
// @Description Create a new todo.
// @Summary create a new todo.
// @Tags Todo
// @Accept json
// @Produce json
// @Param content body string true "Content"
// @Param complete body bool true "Content"
// @Param todo_attrs body models.TodoAttrs true "Todo attributes"
// @Success 200 {object} models.Todo
// @Security ApiKeyAuth
// @Router /api/v1/todos [post]
func CreateTodo(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	todo := &models.Todo{}

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	if err := validate.Struct(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := db.CreateTodo(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"todo":  todo,
	})
}

// UpdateTodo func for updating a todo by given ID.
// @Description Update todo
// @Summary update todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param id body string true "Todo ID"
// @Param content body string true "Content"
// @Param complete body bool true "Complete"
// @Param todo_attrs body models.TodoAttrs true "Todo attributes"
// @Succes 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/todos [put]
func UpdateTodo(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	todo := &models.Todo{}

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundTodo, err := db.GetTodo(todo.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "todo with given ID not found",
		})
	}

	todo.UpdatedAt = time.Now()

	validate := utils.NewValidator()

	if err := validate.Struct(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := db.UpdatedTodo(foundTodo.ID, todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// DeleteTodo func for deleting a todo by given ID.
// @Description Delete todo by given ID.
// @Summary delete todo by given ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id body string true "Todo ID"
// @Success 204 {string} status "no content"
// @Security ApiKeyAuth
// @Router /api/v1/todos [delete]
func DeleteTodo(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	todo := &models.Todo{}

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.StructPartial(todo, "id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundTodo, err := db.GetTodo(todo.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "todo with given ID not found",
		})
	}

	if err := db.DeleteTodo(foundTodo.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
