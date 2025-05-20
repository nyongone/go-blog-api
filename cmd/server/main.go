package main

import (
	"context"
	"fmt"
	"go-blog-api/api/middleware"
	"go-blog-api/api/route"
	"go-blog-api/domain"
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

	defer client.Close()

	app.Use(requestid.New())
	middleware.CORS(app)
	middleware.Logger(app)

	api := app.Group("/api/v1")

	api.Get("/healthCheck", func (ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
			Code: fiber.StatusOK,
		})
	})

	route.NewPostRouter(api.Group("/posts"), context.Background(), client)
	route.NewCategoryRoute(api.Group("/categories"), context.Background(), client)
	route.NewLoginRoute(api.Group("/auth"), context.Background(), client)

	log.Info(app.Listen(fmt.Sprintf("%s:%s", config.EnvVar.AppHost, config.EnvVar.AppPort)))
}