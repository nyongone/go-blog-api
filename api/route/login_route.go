package route

import (
	"context"
	"go-blog-api/api/controller"
	"go-blog-api/ent"
	"go-blog-api/repository"
	"go-blog-api/usecase"

	"github.com/gofiber/fiber/v3"
)

func NewLoginRoute(app fiber.Router, ctx context.Context, client *ent.Client) {
	repository := repository.NewUserRepository(client)
	controller := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(repository),
	}

	app.Post("/login", controller.Login)
}