package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id"` // Primary key ID
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"` // Email with unique constraint
	Password  []byte `json:"-"`                   // Hashed password stored as bytes, excluded from JSON responses for security
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

// Count implements the Entity interface for User
// Returns the total number of user records in the database
func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

// Take implements the Entity interface for User
// Retrieves a paginated subset of users with their roles preloaded
func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User
	db.Preload("Role.Permissions").Offset(offset).Limit(limit).Find(&users)
	return users
}

// SetPassword hashes a plain text password and assigns it to the user
// This encapsulates the password hashing logic within the User model
// Parameters:
//   - password: plain text password to be hashed and stored
//
// Security Note: Uses bcrypt with cost factor 14 for strong security
func (user *User) SetPassword(password string) {
	// Generate a hashed password from the plain text password
	// Cost factor 14 provides strong security (higher = more secure but slower)
	// bcrypt automatically handles salt generation and storage
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	// Assign the hashed password to the user's Password field
	user.Password = hashedPassword
}

// ComparePassword verifies if a provided plain text password matches the stored hash
// This method handles the secure comparison using bcrypt's constant-time comparison
// Parameters:
//   - password: plain text password to verify
//
// Returns:
//   - error: nil if passwords match, error if they don't match or comparison fails
//
// Security: Uses bcrypt.CompareHashAndPassword for secure, timing-attack resistant comparison
func (user *User) ComparePassword(password string) error {
	// Compare the provided password with the stored hash
	// This method handles the cryptographic comparison securely
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
