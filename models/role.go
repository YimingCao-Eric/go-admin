package models

type Role struct {
	Id          uint         `json:"id"` // Primary key ID
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"` // Associated permissions through join table
}
