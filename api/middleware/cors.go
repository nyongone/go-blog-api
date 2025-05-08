package middleware

import (
	"go-blog-api/internal/config"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func CORS(app *fiber.App)  {
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(config.EnvVar.AppCors, ","),
	}))
}