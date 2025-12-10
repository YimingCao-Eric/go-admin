package controllers

import (
	"go-admin/database"
	"go-admin/models"

	"github.com/gofiber/fiber/v3"
)

// AllPermissions retrieves all permissions from the database
// Typically used for populating permission management UI components
func AllPermissions(c fiber.Ctx) error {
	var Permissions []models.Permission

	// Query all permission records
	database.DB.Find(&Permissions)

	return c.JSON(Permissions)
}

// CreatePermission creates a new permission record in the database
// Used to extend the RBAC system with new permission capabilities
// Requires permission name in request body
func CreatePermission(c fiber.Ctx) error {
	var Permission models.Permission

	// Parse JSON request body into Permission struct
	if err := c.Bind().Body(&Permission); err != nil {
		return err
	}

	// Persist new permission to database
	database.DB.Create(&Permission)

	return c.JSON(Permission)
}
