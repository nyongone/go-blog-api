package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Logger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${time}\t${pid}\t${ip}\t${status} - ${method} ${path}\n",
		TimeFormat: "2006/01/06 15:04:05",
		TimeZone: "Asia/Seoul",
	}))
} 