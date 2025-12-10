package middlewares

import (
	"errors"
	"golangProject/database"
	"golangProject/models"
	"golangProject/util"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// IsAuthorized checks if the authenticated user has permission to access a resource
// Implements role-based access control (RBAC) by validating user permissions
// Extracts user from JWT token, loads their role and permissions, then checks access
// Parameters:
//   - page: resource name (e.g., "users", "products") to check permissions for
//
// Permission naming convention:
//   - GET requests require "view_<page>" or "edit_<page>" permission
//   - POST/PUT/DELETE requests require "edit_<page>" permission
//
// Returns error with 401 Unauthorized if user lacks required permission
func IsAuthorized(c fiber.Ctx, page string) error {
	// Extract and validate JWT token
	cookie := c.Cookies("jwt")
	Id, err := util.ParseJWT(cookie)
	if err != nil {
		return err
	}

	// Load user with their role
	userId, _ := strconv.Atoi(Id)
	user := models.User{
		Id: uint(userId),
	}
	database.DB.Preload("Role").Find(&user)

	// Load role with associated permissions
	role := models.Role{
		Id: user.RoleId,
	}
	database.DB.Preload("Permissions").Find(&role)

	// Check permissions based on HTTP method
	if c.Method() == "GET" {
		// GET requests require view or edit permission
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		// POST/PUT/DELETE requests require edit permission
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	// User lacks required permission
	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
