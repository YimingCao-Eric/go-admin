package models

import "gorm.io/gorm"

// Product represents a product in the e-commerce catalog
// Used for managing product inventory and details
type Product struct {
	Id          uint    `json:"id"`          // Primary key
	Title       string  `json:"title"`       // Product name or title
	Description string  `json:"description"` // Detailed product description
	Image       string  `json:"image"`       // Product image URL or file path
	Price       float64 `json:"price"`       // Product price (decimal format)
}

// Count implements the Entity interface for Product
// Returns the total number of product records in the database
// Used by the Paginate function for pagination metadata
func (product *Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Product{}).Count(&total)
	return total
}

// Take implements the Entity interface for Product
// Retrieves a paginated subset of products from the database
// Returns products ordered by their primary key
func (product *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product
	db.Offset(offset).Limit(limit).Find(&products)
	return products
}
