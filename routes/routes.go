package routes

import (
	"golangProject/controllers"
	"golangProject/middlewares"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func Setup(app *fiber.App) {
	// Public routes - no authentication required
	// These endpoints are accessible to unauthenticated users
	app.Post("/api/register", controllers.Register) // Register a new user account
	app.Post("/api/login", controllers.Login)       // Authenticate user and return JWT token

	// Apply authentication middleware to all subsequent routes
	// All routes below this line require a valid JWT token in the request
	app.Use(middlewares.IsAuthenticated)

	// User profile management routes
	// Users can manage their own profile information
	app.Put("/api/users/info", controllers.UpdateInfo)         // Update current user's personal information
	app.Put("/api/users/password", controllers.UpdatePassword) // Change current user's password

	// User session routes
	app.Get("/api/user", controllers.User)      // Get current authenticated user's profile
	app.Post("/api/logout", controllers.Logout) // Invalidate user session and logout

	// User management routes (admin operations)
	// Full CRUD operations for user management
	app.Get("/api/users", controllers.AllUsers)          // Retrieve paginated list of all users
	app.Post("/api/users", controllers.CreateUser)       // Create a new user account
	app.Get("/api/users/:id", controllers.GetUser)       // Retrieve user details by ID
	app.Put("/api/users/:id", controllers.UpdateUser)    // Update user information by ID
	app.Delete("/api/users/:id", controllers.DeleteUser) // Delete a user account by ID

	// Role management routes
	// Role-based access control (RBAC) operations
	app.Get("/api/roles", controllers.AllRoles)          // Retrieve list of all roles
	app.Post("/api/roles", controllers.CreateRole)       // Create a new role
	app.Get("/api/roles/:id", controllers.GetRole)       // Retrieve role details by ID
	app.Put("/api/roles/:id", controllers.UpdateRole)    // Update role information by ID
	app.Delete("/api/roles/:id", controllers.DeleteRole) // Delete a role by ID

	// Permission management routes
	app.Get("/api/permissions", controllers.AllPermissions)    // Retrieve list of all permissions
	app.Post("/api/permissions", controllers.CreatePermission) // Create a new permission

	// Product management routes
	// Full CRUD operations for product catalog
	app.Get("/api/products", controllers.AllProducts)          // Retrieve paginated list of products
	app.Post("/api/products", controllers.CreateProduct)       // Create a new product
	app.Get("/api/products/:id", controllers.GetProduct)       // Retrieve product details by ID
	app.Put("/api/products/:id", controllers.UpdateProduct)    // Update product information by ID
	app.Delete("/api/products/:id", controllers.DeleteProduct) // Delete a product by ID

	// File upload and serving routes
	app.Post("/api/upload", controllers.Upload)       // Upload files via multipart form data
	app.Get("/api/uploads*", static.New("./uploads")) // Serve uploaded files as static content

	// Order management and analytics routes
	app.Get("/api/orders", controllers.AllOrders) // Retrieve paginated orders with associated items
	app.Post("/api/export", controllers.Export)   // Export orders data to CSV format
	app.Get("/api/chart", controllers.Chart)      // Retrieve sales analytics data for chart visualization
}
