package controller

import (
	"go-blog-api/domain"
	"go-blog-api/ent"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryUsecase
}

func (c *CategoryController) GetCategory(ctx fiber.Ctx) error {
	categoryId, err := strconv.Atoi(ctx.Params("categoryId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	category, err := c.CategoryUsecase.GetCategory(ctx.Context(), categoryId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
		Result: category,
	})
}

func (c *CategoryController) GetCategoryList(ctx fiber.Ctx) error {
	category, _ := c.CategoryUsecase.GetCategoryList(ctx.Context())

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
		Result: category,
	})
}

func (c *CategoryController) CreateCategory(ctx fiber.Ctx) error {
	var category *ent.Category

	err := ctx.Bind().Body(&category)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	err = c.CategoryUsecase.CreateCategory(ctx.Context(), category)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&domain.ErrorResponse{
		Code: fiber.StatusCreated,
	})
}

func (c *CategoryController) UpdateCategory(ctx fiber.Ctx) error {
	var category *ent.Category

	categoryId, err := strconv.Atoi(ctx.Params("categoryId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	_, err = c.CategoryUsecase.GetCategory(ctx.Context(), categoryId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	err = ctx.Bind().Body(&category)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	err = c.CategoryUsecase.UpdateCategory(ctx.Context(), categoryId, category)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
	})
}

func (c *CategoryController) DeleteCategory(ctx fiber.Ctx) error {
	categoryId, err := strconv.Atoi(ctx.Params("categoryId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	_, err = c.CategoryUsecase.GetCategory(ctx.Context(), categoryId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.ErrorResponse{
			Code: fiber.StatusNotFound,
		})
	}

	err = c.CategoryUsecase.DeleteCategory(ctx.Context(), categoryId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.ErrorResponse{
		Code: fiber.StatusOK,
	})
}