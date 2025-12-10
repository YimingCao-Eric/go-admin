package models

import "gorm.io/gorm"

// Entity defines the interface for models that support pagination
// Models implementing this interface can use the generic Paginate function
// This provides a consistent pagination pattern across different entity types
type Entity interface {
	// Count returns the total number of records in the database for this entity type
	Count(db *gorm.DB) int64

	// Take retrieves a paginated subset of records from the database
	// Parameters:
	//   - limit: maximum number of records to return per page
	//   - offset: number of records to skip (for pagination)
	// Returns: slice of entity instances
	Take(db *gorm.DB, limit int, offset int) interface{}
}
