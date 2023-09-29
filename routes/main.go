package routes

import (
	"go-logger/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	requestLogs := app.Group("/request_log")
	requestLogs.Get("/", handlers.HandleGetAllRequestLogs)
	requestLogs.Post("/", handlers.HandleCreateRequestLog)
}
