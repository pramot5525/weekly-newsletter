package router

import (
	"weekly-newsletter/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app     *fiber.App
	handler *handler.Handler
}

func NewRouter(app *fiber.App, handler *handler.Handler) error {
	router := Router{
		app:     app,
		handler: handler,
	}
	router.initRouter()
	return nil
}

func (router *Router) initRouter() {
	app := router.app
	handler := router.handler

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
	})

	api := app.Group("/api/v1")
	user := api.Group("/user")
	{
		user.Post("/subscribe", handler.Subscribe)
		user.Get("/unsubscribe", handler.Unsubscribe)
		user.Get("/publish", handler.Publish)
	}
}
