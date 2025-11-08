package controllers

import (
	"golangProject/database"
	"golangProject/models"

	"github.com/gofiber/fiber/v3"
)

// AllPermissions retrieves all permissions from the database
// This is typically used for permission management UI
func AllPermissions(c fiber.Ctx) error {
	// Create a slice to hold all permission records
	var Permissions []models.Permission

	// Query the database to find all permissions
	// This executes: SELECT * FROM permissions;
	database.DB.Find(&Permissions)

	// Return the permissions as JSON response
	return c.JSON(Permissions)
}

// CreatePermission creates a new permission in the database
// This function allows creating new permissions for the RBAC system
func CreatePermission(c fiber.Ctx) error {
	// Create a Permission struct to hold the request data
	var Permission models.Permission

	// Parse the JSON request body into the permission struct
	if err := c.Bind().Body(&Permission); err != nil {
		return err
	}

	// Create the new permission record in the database
	// This executes: INSERT INTO permissions (name) VALUES (?);
	database.DB.Create(&Permission)

	// Return the created permission as JSON response
	return c.JSON(Permission)
}
