package models

// Permission represents a permission in the role-based access control (RBAC) system
// Permissions define granular access rights that can be assigned to roles
// Examples: "view_users", "edit_products", "delete_orders"
type Permission struct {
	Id   uint   `json:"id"`   // Primary key
	Name string `json:"name"` // Permission identifier (e.g., "view_users", "edit_products")
}
