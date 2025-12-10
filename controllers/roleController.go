package controllers

import (
	"golangProject/database"
	"golangProject/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// AllRoles retrieves all roles with their associated permissions
// Typically used for role management UI to display available roles
func AllRoles(c fiber.Ctx) error {
	var roles []models.Role

	// Load all roles with preloaded permissions
	database.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

// CreateRole creates a new role with associated permissions
// Establishes a many-to-many relationship between roles and permissions
// Request body: { "name": string, "permissions": []string (permission IDs) }
func CreateRole(c fiber.Ctx) error {
	var roleDTO fiber.Map

	// Parse JSON request body
	if err := c.Bind().Body(&roleDTO); err != nil {
		return err
	}

	// Extract permissions array from request
	list := roleDTO["permissions"].([]interface{})

	// Build Permission slice from provided IDs
	permissions := make([]models.Permission, len(list))
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	// Create role with associated permissions
	role := models.Role{
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	// Persist role and create associations in join table
	database.DB.Create(&role)

	return c.JSON(role)
}

// GetRole retrieves a specific role by ID with its associated permissions
// Used for viewing role details and permission assignments
// URL parameter: id (role identifier)
func GetRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	// Find role and eagerly load permissions
	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

// RolePermission represents the role_permissions join table structure
// Used for direct manipulation of the many-to-many relationship
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"` // Foreign key to roles table
	PermissionID uint `gorm:"primaryKey"` // Foreign key to permissions table
}

// UpdateRole updates an existing role's name and permission assignments
// Replaces all existing permission associations with the new set
// URL parameter: id (role identifier to update)
func UpdateRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map

	// Parse JSON request body
	if err := c.Bind().Body(&roleDTO); err != nil {
		return err
	}

	// Extract and process permissions array
	list := roleDTO["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	// Remove all existing permission associations
	var rolePermission RolePermission
	database.DB.Table("role_permissions").Where("role_id = ?", id).Delete(&rolePermission)

	// Update role with new name and permissions
	role := models.Role{
		Id:          uint(id),
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}
	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

// DeleteRole permanently removes a role from the database
// Cascades deletion to role_permissions join table associations
// This is a destructive operation - ensure proper authorization
// URL parameter: id (role identifier to delete)
func DeleteRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	// Delete role (GORM handles join table cleanup automatically)
	database.DB.Delete(&role)

	return nil
}
