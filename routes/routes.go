package routes

import (
	"golangProject/controllers"
	"golangProject/middlewares"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func Setup(app *fiber.App) {
	// Public routes - accessible without authentication
	app.Post("/api/register", controllers.Register) // User registration endpoint
	app.Post("/api/login", controllers.Login)       // User login endpoint

	// Apply authentication middleware to all routes below this line
	app.Use(middlewares.IsAuthenticated)

	// User profile management routes - users can manage their own data
	app.Put("/api/users/info", controllers.UpdateInfo)         // Update personal information
	app.Put("/api/users/password", controllers.UpdatePassword) // Change password

	// Protected routes - require valid JWT token
	app.Get("/api/user", controllers.User)      // Get current user profile
	app.Post("/api/logout", controllers.Logout) // User logout endpoint

	// User management routes (typically admin functions)
	app.Get("/api/users", controllers.AllUsers)          // Get all users (list view)
	app.Post("/api/users", controllers.CreateUser)       // Create a new user programmatically
	app.Get("/api/users/:id", controllers.GetUser)       // Get a specific user by ID
	app.Put("/api/users/:id", controllers.UpdateUser)    // Update a specific user by ID
	app.Delete("/api/users/:id", controllers.DeleteUser) // Delete a specific user by ID

	// Role-based access control (RBAC) routes
	app.Get("/api/roles", controllers.AllRoles)          // Get all roles (list view)
	app.Post("/api/roles", controllers.CreateRole)       // Create a new role programmatically
	app.Get("/api/roles/:id", controllers.GetRole)       // Get a specific role by ID
	app.Put("/api/roles/:id", controllers.UpdateRole)    // Update a specific role by ID
	app.Delete("/api/roles/:id", controllers.DeleteRole) // Delete a specific role by ID

	// Permission management routes
	app.Get("/api/permissions", controllers.AllPermissions)    // Get all permission (list view)
	app.Post("/api/permissions", controllers.CreatePermission) // Create a new permission programmatically

	// Product management routes
	app.Get("/api/products", controllers.AllProducts)          // Get paginated list of products
	app.Post("/api/products", controllers.CreateProduct)       // Create a new product
	app.Get("/api/products/:id", controllers.GetProduct)       // Get a specific product by ID
	app.Put("/api/products/:id", controllers.UpdateProduct)    // Update a specific product by ID
	app.Delete("/api/products/:id", controllers.DeleteProduct) // Delete a specific product by ID

	// File upload route - handles receiving files from clients via multipart form data
	app.Post("/api/upload", controllers.Upload)       // Clients can POST files to this endpoint for storage on the server
	app.Get("/api/uploads*", static.New("./uploads")) // Static file serving route - makes uploaded files publicly accessible

	// Order management and analytics routes
	app.Get("/api/orders", controllers.AllOrders) // Get paginated orders with items
	app.Post("/api/export", controllers.Export)   // Export orders to CSV file
	app.Get("/api/chart", controllers.Chart)      // Get sales data for charts
}
