package config

import "github.com/gofiber/fiber/v3"

func FiberConfig() fiber.Config {
	return fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader: "BLOG",
		AppName: "Go Blog API",
	}
}