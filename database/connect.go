package database

import (
	"go-admin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB holds the global database connection instance
// This connection pool is shared across the entire application for database operations
var DB *gorm.DB

// Connect establishes a connection to the MySQL database and performs auto-migration
// Connection string format: "username:password@/database_name"
// Automatically migrates all application models to ensure database schema matches code
// Panics if database connection cannot be established
func Connect() {
	// Establish connection to MySQL database
	// Note: Connection credentials are hardcoded - consider using environment variables for production
	db, err := gorm.Open(mysql.Open("root:fb112358@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Store connection in global variable for application-wide access
	DB = db

	// Auto-migrate database schema for all models
	// Creates tables if they don't exist and updates schema for existing tables
	// Models included: User, Role, Permission, Product, Order, OrderItem
	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
}
