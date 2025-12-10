package controllers

import (
	"golangProject/database"
	"golangProject/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// AllProducts retrieves a paginated list of products from the database
// Uses the generic Paginate function for consistent pagination response format
// Query parameter: page (defaults to 1 if not provided)
func AllProducts(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

// CreateProduct creates a new product record in the database
// Adds a new item to the product catalog
// Request body should contain: title, description, image, price
func CreateProduct(c fiber.Ctx) error {
	var product models.Product

	// Parse JSON request body
	if err := c.Bind().Body(&product); err != nil {
		return err
	}

	// Persist new product to database
	database.DB.Create(&product)

	return c.JSON(product)
}

// GetProduct retrieves a specific product by ID
// Used for viewing individual product details
// URL parameter: id (product identifier)
func GetProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	// Find product by primary key
	database.DB.Find(&product)

	return c.JSON(product)
}

// UpdateProduct updates an existing product's information
// Allows modification of: title, description, image, and price
// URL parameter: id (product identifier to update)
func UpdateProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	// Parse updated product data from request body
	if err := c.Bind().Body(&product); err != nil {
		return err
	}

	// Update product record in database
	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

// DeleteProduct permanently removes a product from the database
// This is a destructive operation - ensure proper authorization is in place
// URL parameter: id (product identifier to delete)
func DeleteProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	// Delete product record from database
	database.DB.Delete(&product)

	return nil
}
