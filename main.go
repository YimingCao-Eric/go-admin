package main

import (
	"golangProject/database"
	"golangProject/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	// Connect to the database
	database.Connect()

	// Initialize a new Fiber app
	app := fiber.New()

	// Configure and use CORS middleware
	app.Use(cors.New(cors.Config{
		// Custom function to check allowed origins
		AllowOriginsFunc: func(origin string) bool {

			return origin == "http://localhost:3000"
		},
		AllowCredentials: true, // Allow cookies and credentials to be sent
	}))

	// Set up application routes
	routes.Setup(app)

	// Start the server on port 3000
	app.Listen(":8000")
}
