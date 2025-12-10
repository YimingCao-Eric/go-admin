package main

import (
	"go-admin/database"
	"go-admin/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	// Establish database connection
	database.Connect()

	// Create a new Fiber application instance
	app := fiber.New()

	// Configure Cross-Origin Resource Sharing (CORS) middleware
	app.Use(cors.New(cors.Config{
		// Currently allows requests from localhost:3000 (typical frontend dev server)
		AllowOriginsFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		// AllowCredentials enables cookies and authorization headers in CORS requests
		AllowCredentials: true,
	}))

	// This sets up all API endpoints(routes) and their corresponding handlers
	routes.Setup(app)

	// Start the HTTP server and listen on port 8000
	app.Listen(":8000")
}
