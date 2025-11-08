package controllers

import (
	"golangProject/database"
	"golangProject/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// AllRoles retrieves all roles from the database
// This is typically an admin-only function to view all available roles
func AllRoles(c fiber.Ctx) error {
	// Create a slice to hold all role records
	var roles []models.Role

	// Query the database to find all roles
	// This executes: SELECT * FROM roles;
	database.DB.Preload("Permissions").Find(&roles)

	// Return the roles as JSON response
	return c.JSON(roles)
}

// CreateRole creates a new role with associated permissions
// This function allows creating roles programmatically with permission assignments
func CreateRole(c fiber.Ctx) error {
	// Create a map to hold the request data (name and permissions array)
	var roleDTO fiber.Map

	// Parse the JSON request body into the map
	if err := c.Bind().Body(&roleDTO); err != nil {
		return err
	}

	// Extract the permissions array from the request data
	// Type assertion to convert interface{} to []interface{}
	list := roleDTO["permissions"].([]interface{})

	// Create a slice of Permission models with the provided IDs
	permissions := make([]models.Permission, len(list))

	// Convert each permission ID from string to uint and create Permission objects
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	// Create the new Role instance with name and associated permissions
	role := models.Role{
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	// Create the new role record in the database
	// This also creates the associations in the role_permissions join table
	database.DB.Create(&role)

	// Return the created role as JSON response
	return c.JSON(role)
}

// GetRole retrieves a specific role by ID from the database including its permissions
// This is used for viewing individual role details with associated permissions
func GetRole(c fiber.Ctx) error {
	// Extract the role ID from the URL parameter and convert to integer
	id, _ := strconv.Atoi(c.Params("id"))

	// Create a Role instance with the ID set for the database query
	role := models.Role{
		Id: uint(id),
	}

	// Find the role in the database by primary key and preload its permissions
	// Preload eagerly loads the associated permissions from the join table
	// This executes: SELECT * FROM roles WHERE id = ?;
	// And: SELECT * FROM permissions JOIN role_permissions ON ...
	database.DB.Preload("Permissions").Find(&role)

	// Return the role with permissions as JSON response
	return c.JSON(role)
}

// RolePermission represents the join table for role-permission many-to-many relationship
// This struct is used for direct operations on the role_permissions table
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}

// UpdateRole updates an existing role's information and permissions
// This allows modifying role name and changing permission assignments
func UpdateRole(c fiber.Ctx) error {
	// Extract the role ID from the URL parameter and convert to integer
	id, _ := strconv.Atoi(c.Params("id"))

	// Create a map to hold the updated role data
	var roleDTO fiber.Map

	// Parse the JSON request body into the map
	if err := c.Bind().Body(&roleDTO); err != nil {
		return err
	}

	// Extract the permissions array from the request data
	list := roleDTO["permissions"].([]interface{})

	// Create a slice of Permission models with the updated permission IDs
	permissions := make([]models.Permission, len(list))

	// Convert each permission ID and create Permission objects
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	// Define a struct to represent the role_permissions join table
	// This struct maps to the join table with composite primary keys
	var rolePermission RolePermission

	// Delete all existing permission associations for this role from the join table
	// This executes: DELETE FROM role_permissions WHERE role_id = ?
	database.DB.Table("role_permissions").Where("role_id = ?", id).Delete(&rolePermission)

	// Create the Role instance with updated data
	role := models.Role{
		Id:          uint(id),
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	// Update the role record in the database
	// This updates the role name and manages the permission associations
	database.DB.Model(&role).Updates(role)

	// Return the updated role as JSON response
	return c.JSON(role)
}

// DeleteRole removes a role from the database
// This is a destructive operation and should be protected with proper authorization
func DeleteRole(c fiber.Ctx) error {
	// Extract the role ID from the URL parameter and convert to integer
	id, _ := strconv.Atoi(c.Params("id"))

	// Create a Role instance with the target ID for deletion
	role := models.Role{
		Id: uint(id),
	}

	// Delete the role record from the database
	// This automatically handles the removal from role_permissions join table due to GORM associations
	database.DB.Delete(&role)

	// Return nil (no content) to indicate successful deletion
	return nil
}
