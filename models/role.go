package models

// Role represents a user role in the role-based access control (RBAC) system
// Roles group multiple permissions together and are assigned to users
// Maintains a many-to-many relationship with Permissions via the role_permissions join table
type Role struct {
	Id          uint         `json:"id"`          // Primary key
	Name        string       `json:"name"`        // Role name (e.g., "admin", "editor", "viewer")
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"` // Associated permissions
}
