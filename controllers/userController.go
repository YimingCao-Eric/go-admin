package controllers

import (
	"go-admin/database"
	"go-admin/middlewares"
	"go-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// AllUsers retrieves a paginated list of users from the database
// Requires authorization with "users" permission
// Uses the generic Paginate function for consistent pagination response format
// Query parameter: page (defaults to 1 if not provided)
func AllUsers(c fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

// CreateUser creates a new user account programmatically
// Requires authorization with "users" permission (admin function)
// Sets a default password - in production, consider password reset flow
// Request body should contain: first_name, last_name, email, role_id
func CreateUser(c fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var user models.User

	// Parse JSON request body
	if err := c.Bind().Body(&user); err != nil {
		return err
	}

	// Set default password (should be changed by user on first login)
	user.SetPassword("3")

	// Persist new user to database
	database.DB.Create(&user)

	return c.JSON(user)
}

// GetUser retrieves a specific user by ID with their role information
// Requires authorization with "users" permission
// Used for viewing individual user profiles
// URL parameter: id (user identifier)
func GetUser(c fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	// Find user with preloaded role information
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

// UpdateUser updates an existing user's information
// Requires authorization with "users" permission
// Allows modification of: first_name, last_name, email
// URL parameter: id (user identifier to update)
func UpdateUser(c fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	// Parse updated user data from request body
	if err := c.Bind().Body(&user); err != nil {
		return err
	}

	// Update user record in database
	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

// DeleteUser permanently removes a user from the database
// Requires authorization with "users" permission
// This is a destructive operation - ensure proper authorization is in place
// URL parameter: id (user identifier to delete)
func DeleteUser(c fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	// Delete user record from database
	database.DB.Delete(&user)

	return nil
}
