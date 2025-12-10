package models

import (
	"math"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// Paginate provides a generic pagination solution for any Entity type
// Eliminates code duplication by providing a consistent pagination pattern
// Works with any model that implements the Entity interface
//
// Parameters:
//   - db: GORM database connection instance
//   - entity: Entity interface implementation (User, Product, Order, etc.)
//   - page: current page number (1-indexed)
//
// Returns:
//   - fiber.Map containing:
//     - "data": paginated records from the entity
//     - "meta": pagination metadata (total, page, last_page)
func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	// Records per page (configurable - currently set to 5)
	limit := 5

	// Calculate offset for SQL LIMIT/OFFSET query
	offset := (page - 1) * limit

	// Retrieve paginated data using entity's Take method
	data := entity.Take(db, limit, offset)

	// Get total record count using entity's Count method
	total := entity.Count(db)

	// Return standardized pagination response
	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)), // Calculate total number of pages
		},
	}
}
