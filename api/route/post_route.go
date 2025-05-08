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

func NewPostRouter(app fiber.Router, ctx context.Context, client *ent.Client) {
	repository := repository.NewPostRepository(client)
	controller := &controller.PostController{
		PostUsecase: usecase.NewPostUsecase(repository),
	}

	app.Get("", controller.GetPostList)
	app.Get(":postId", controller.GetPost)
	app.Get("slug/:postSlug", controller.GetPostBySlug)
	app.Post("", controller.CreatePost, middleware.AuthMiddleware)
	app.Put(":postId", controller.UpdatePost, middleware.AuthMiddleware)
	app.Delete(":postId", controller.DeletePost, middleware.AuthMiddleware)
}