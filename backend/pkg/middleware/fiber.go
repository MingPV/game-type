package middleware

import (
	"github.com/MingPV/clean-go-template/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// LoadCommon sets common global middleware for the app
func FiberMiddleware(app *fiber.App, cfg *config.Config) {
	app.Use(

		logger.New(), // Logs all requests

		cors.New(cors.Config{
			AllowOrigins:     cfg.FrontendURL, // need to be changed in production
			AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowCredentials: true,
		}),
	)
}
