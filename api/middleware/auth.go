package middleware

import (
	"go-blog-api/domain"
	"go-blog-api/internal/util"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(ctx fiber.Ctx) error {
	authHeader := strings.Split(ctx.Get("Authorization"), " ")
	if len(authHeader) == 2 {
		accessToken := authHeader[1]
		authorized, err := util.VerifyToken(accessToken)
		if !authorized || err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&domain.ErrorResponse{
				Code: fiber.StatusUnauthorized,
			})
		}

		return ctx.Next()
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(&domain.ErrorResponse{
		Code: fiber.StatusUnauthorized,
	})
}