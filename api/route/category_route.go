package route

import (
	"context"
	"go-blog-api/api/controller"
	"go-blog-api/api/middleware"
	"go-blog-api/ent"
	"go-blog-api/repository"
	"go-blog-api/usecase"

	"github.com/gofiber/fiber/v3"
)

func NewCategoryRoute(app fiber.Router, ctx context.Context, client *ent.Client) {
	repository := repository.NewCategoryRepository(client)
	controller := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(repository),
	}

	app.Get("", controller.GetCategoryList)
	app.Get(":categoryId", controller.GetCategory)
	app.Post("", controller.CreateCategory, middleware.AuthMiddleware)
	app.Put(":categoryId", controller.UpdateCategory, middleware.AuthMiddleware)
	app.Delete(":categoryId", controller.DeleteCategory, middleware.AuthMiddleware)
}