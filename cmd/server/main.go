package main

import (
	"context"
	"go-blog-api/api/middleware"
	"go-blog-api/api/route"
	"go-blog-api/internal/config"
	"go-blog-api/internal/datastore"

	"github.com/gofiber/fiber/v3/log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func main() {
	config.LoadEnv()

	app := fiber.New(config.FiberConfig())

	client, err := datastore.NewClient()

	if err != nil {
		log.Fatalf("Failed connecting mysql database: %v", err)
	}

	app.Use(requestid.New())
	middleware.CORS(app)
	middleware.Logger(app)

	api := app.Group("/api")

	route.NewPostRouter(api.Group("/v1/posts"), context.Background(), client)
	route.NewCategoryRoute(api.Group("/v1/categories"), context.Background(), client)
	route.NewLoginRoute(api.Group("/v1/auth"), context.Background(), client)

	log.Info(app.Listen(":3000"))
}