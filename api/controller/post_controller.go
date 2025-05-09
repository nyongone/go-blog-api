package controller

import (
	"go-blog-api/domain"
	"go-blog-api/ent"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type PostController struct {
	PostUsecase domain.PostUsecase
}

func (c *PostController) GetPost(ctx fiber.Ctx) error {
	postId, err := strconv.Atoi(ctx.Params("postId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	post, err := c.PostUsecase.GetPost(ctx.Context(), postId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
		Result: post,
	})
}

func (c *PostController) GetPostBySlug(ctx fiber.Ctx) error {
	postSlug := ctx.Params("postSlug")

	if postSlug == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	post, err := c.PostUsecase.GetPostBySlug(ctx.Context(), postSlug)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
		Result: post,
	})
}

func (c *PostController) GetPostList(ctx fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page"))

	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))

	if err != nil || limit <= 0 {
		limit = 10
	}

	posts, _ := c.PostUsecase.GetPostList(ctx.Context(), page, limit)

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
		Result: posts,
	})
}

func (c *PostController) CreatePost(ctx fiber.Ctx) error {
	var post *ent.Post

	err := ctx.Bind().Body(&post)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	err = c.PostUsecase.CreatePost(ctx.Context(), post)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&domain.SuccessResponse{
		Code: fiber.StatusCreated,
	})
}

func (c *PostController) UpdatePost(ctx fiber.Ctx) error {
	var post *ent.Post

	postId, err := strconv.Atoi(ctx.Params("postId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.SuccessResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	_, err = c.PostUsecase.GetPost(ctx.Context(), postId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	err = ctx.Bind().Body(&post)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	err = c.PostUsecase.UpdatePost(ctx.Context(), postId, post)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
	})
}

func (c *PostController) DeletePost(ctx fiber.Ctx) error {
	postId, err := strconv.Atoi(ctx.Params("postId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	_, err = c.PostUsecase.GetPost(ctx.Context(), postId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	err = c.PostUsecase.DeletePost(ctx.Context(), postId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
	})
}