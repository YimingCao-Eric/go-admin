package database

import (
	"golangProject/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB Global variable to hold database connection
var DB *gorm.DB

func Connect() {
	// Open connection to MySQL database with connection string
	db, err := gorm.Open(mysql.Open("root:fb112358@/go_admin"), &gorm.Config{})

	// Check if connection failed
	if err != nil {
		panic("failed to connect database")
	}

	// Assign the database connection to global variable
	DB = db

	// Auto-migrate all application models to ensure database schema is up to date
	// Creates or updates tables for: users, roles, permissions, products, orders, order_items
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})

}
