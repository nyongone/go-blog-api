package controller

import (
	"go-blog-api/domain"

	"github.com/gofiber/fiber/v3"
)

type LoginController struct {
	LoginUsecase	domain.LoginUsecase
}

func (c *LoginController) Login(ctx fiber.Ctx) error {
	var request *domain.LoginRequest

	err := ctx.Bind().Body(&request)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	user, err := c.LoginUsecase.GetUserByEmail(ctx.Context(), request.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	// TODO: Encrypting Password
	if user.Password != request.Password {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	accessToken, err := c.LoginUsecase.CreateAccessToken(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	refreshToken, err := c.LoginUsecase.CreateRefreshToken()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	err = c.LoginUsecase.UpdateRefreshToken(ctx.Context(), user.ID, refreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.ErrorResponse{
			Code: fiber.StatusBadRequest,
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name: "access_token",
		Value: accessToken,
	})

	ctx.Cookie(&fiber.Cookie{
		Name: "refresh_token",
		Value: refreshToken,
	})
	
	return ctx.Status(fiber.StatusOK).JSON(&domain.SuccessResponse{
		Code: fiber.StatusOK,
	})
}