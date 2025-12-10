package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user account in the system
// Maintains authentication credentials and role-based access control
type User struct {
	Id        uint   `json:"id"`         // Primary key
	FirstName string `json:"first_name"` // User's first name
	LastName  string `json:"last_name"`  // User's last name
	Email     string `json:"email" gorm:"unique"` // Email address (unique constraint)
	Password  []byte `json:"-"`          // Hashed password (excluded from JSON for security)
	RoleId    uint   `json:"role_id"`    // Foreign key to Role
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"` // Associated role
}

// Count implements the Entity interface for User
// Returns the total number of user records in the database
// Used by the Paginate function for pagination metadata
func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

// Take implements the Entity interface for User
// Retrieves a paginated subset of users with their roles and permissions preloaded
// Eagerly loads Role.Permissions to avoid N+1 query problem
func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User
	db.Preload("Role.Permissions").Offset(offset).Limit(limit).Find(&users)
	return users
}

// SetPassword hashes a plain text password using bcrypt and stores it
// Encapsulates password hashing logic within the User model
// Uses bcrypt cost factor 14 for strong security (higher = more secure but slower)
// Automatically generates and stores a unique salt for each password
//
// Parameters:
//   - password: plain text password to be hashed
func (user *User) SetPassword(password string) {
	// Generate bcrypt hash with cost factor 14
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

// ComparePassword verifies if a provided plain text password matches the stored hash
// Uses bcrypt's constant-time comparison to prevent timing attacks
// Returns nil if passwords match, error otherwise
//
// Parameters:
//   - password: plain text password to verify
//
// Returns:
//   - error: nil if passwords match, error if mismatch or comparison fails
func (user *User) ComparePassword(password string) error {
	// Compare provided password with stored hash using bcrypt
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
