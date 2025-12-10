package middlewares

import (
	"golangProject/util"

	"github.com/gofiber/fiber/v3"
)

// IsAuthenticated validates JWT token presence and authenticity
// Protects routes that require user authentication
// Extracts token from "jwt" cookie and verifies its validity
// Returns 401 Unauthorized if token is missing or invalid
// Usage: app.Use(middlewares.IsAuthenticated) to protect all routes below,
//        or app.Get("/protected", middlewares.IsAuthenticated, handler) for specific routes
func IsAuthenticated(c fiber.Ctx) error {
	// Extract JWT token from authentication cookie
	cookie := c.Cookies("jwt")

	// Validate token using utility function
	if _, err := util.ParseJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	// Token is valid - proceed to next handler
	return c.Next()
}
