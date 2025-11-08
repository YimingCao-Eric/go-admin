package models

import "gorm.io/gorm"

// Order represents a customer order in the e-commerce system
// Contains customer information and links to order items
type Order struct {
	Id         uint        `json:"id"` // Primary key, unique order identifier
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"` // Computed full name (FirstName + LastName), not stored in database
	Email      string      `json:"email"`
	Total      float32     `json:"total" gorm:"-"` // Computed order total sum, not stored in database
	UpdateAt   string      `json:"update_at"`
	CreateAt   string      `json:"create_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"` // Associated order items, one-to-many relationship
}

// OrderItem represents an individual product within an order
// Each order can contain multiple order items (products)
type OrderItem struct {
	Id           uint    `json:"id"`       // Primary key
	OrderId      uint    `json:"order_id"` // Foreign key linking to the parent order
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

// Count implements the Entity interface for Order
// Returns the total number of orders in the database
// Used by the generic Paginate function for pagination metadata
func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

// Take implements the Entity interface for Order
// Retrieves a paginated subset of orders with computed fields
// This method:
//   - Preloads associated OrderItems to avoid N+1 query problem
//   - Calculates the total price for each order
//   - Computes the full customer name from first/last name
func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	// Retrieve paginated orders with their associated items
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	// Process each order to compute derived fields
	for i, _ := range orders {
		var total float32 = 0

		// Calculate order total by summing all item prices * quantities
		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}

		// Compute full name from first and last name
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName

		// Store computed total (note: slight floating point precision may occur)
		orders[i].Total = total
	}
	return orders
}
