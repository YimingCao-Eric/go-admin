package models

import "gorm.io/gorm"

// Order represents a customer order in the e-commerce system
// Contains customer information and maintains a one-to-many relationship with OrderItems
type Order struct {
	Id         uint        `json:"id"`                           // Primary key, unique order identifier
	FirstName  string      `json:"-"`                            // Customer first name (not included in JSON response)
	LastName   string      `json:"-"`                            // Customer last name (not included in JSON response)
	Name       string      `json:"name" gorm:"-"`                // Computed full name (FirstName + LastName), virtual field
	Email      string      `json:"email"`                        // Customer email address
	Total      float32     `json:"total" gorm:"-"`               // Computed order total, virtual field
	UpdateAt   string      `json:"update_at"`                    // Last update timestamp
	CreateAt   string      `json:"create_at"`                    // Creation timestamp
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"` // Associated order items
}

// OrderItem represents an individual product within an order
// Each order can contain multiple order items, forming a one-to-many relationship
type OrderItem struct {
	Id           uint    `json:"id"`            // Primary key
	OrderId      uint    `json:"order_id"`      // Foreign key to parent Order
	ProductTitle string  `json:"product_title"` // Product name at time of purchase
	Price        float32 `json:"price"`         // Product price at time of purchase
	Quantity     uint    `json:"quantity"`      // Quantity of this product in the order
}

// Count implements the Entity interface for Order
// Returns the total number of order records in the database
// Used by the Paginate function to calculate pagination metadata
func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

// Take implements the Entity interface for Order
// Retrieves a paginated subset of orders and computes derived fields
// Preloads OrderItems to avoid N+1 query problem
// Calculates order total and full customer name for each order
func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	// Retrieve paginated orders with eagerly loaded order items
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	// Compute derived fields for each order
	for i := range orders {
		var total float32 = 0

		// Calculate order total by summing (price * quantity) for all items
		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}

		// Compute full customer name
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		orders[i].Total = total
	}
	return orders
}
