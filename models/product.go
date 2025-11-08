package models

import "gorm.io/gorm"

type Product struct {
	Id          uint    `json:"id"`          // Primary key ID for the product
	Title       string  `json:"title"`       // Product name or title
	Description string  `json:"description"` // Detailed description of the product
	Image       string  `json:"image"`       // Product image URL or path
	Price       float64 `json:"price"`       // Product price in decimal format
}

// Count implements the Entity interface for Product
// Returns the total number of product records in the database
func (product *Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Product{}).Count(&total)
	return total
}

// Take implements the Entity interface for Product
// Retrieves a paginated subset of products from the database
func (product *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product
	db.Offset(offset).Limit(limit).Find(&products)
	return products
}
