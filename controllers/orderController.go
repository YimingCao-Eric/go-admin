package controllers

import (
	"encoding/csv"
	"go-admin/database"
	"go-admin/models"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// AllOrders retrieves a paginated list of all orders with their associated items
// Uses the generic Paginate function for consistent pagination response format
// Query parameter: page (defaults to 1 if not provided)
func AllOrders(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}

// Export generates a CSV file containing all orders and order items
// Creates a structured export file suitable for spreadsheet applications
// The CSV file is generated and sent to the client as a download
func Export(c fiber.Ctx) error {
	filePath := "./csv/order.csv"

	// Generate CSV file with order data
	if err := CreateFile(filePath); err != nil {
		return err
	}

	// Send file to client as download
	return c.Download(filePath)
}

// CreateFile generates a CSV file with order and order item data
// CSV structure: each order has one header row with customer info,
// followed by one row per order item (empty cells for customer columns)
// This format allows visual grouping of items under their parent order
func CreateFile(filePath string) error {
	// Create CSV file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialize CSV writer with automatic flushing
	writer := csv.NewWriter(file)
	defer writer.Flush()

	var orders []models.Order

	// Load all orders with preloaded order items
	database.DB.Preload("OrderItems").Find(&orders)

	// Write CSV header row
	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	// Write order data to CSV
	for _, order := range orders {
		// Write order header row with customer information
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"",
			"",
			"",
		}
		if err := writer.Write(data); err != nil {
			return err
		}

		// Write each order item as a separate row
		for _, orderItem := range order.OrderItems {
			data := []string{
				"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),
			}
			if err := writer.Write(data); err != nil {
				return err
			}
		}
	}
	return nil
}

// Sales represents daily sales totals for chart visualization
// Used by the Chart endpoint to provide sales trend data over time
type Sales struct {
	Date string `json:"date"` // Date in YYYY-MM-DD format
	Sum  string `json:"sum"`  // Total sales amount for that date
}

// Chart retrieves daily sales data aggregated by date
// Uses raw SQL to group orders by creation date and calculate daily totals
// Returns data formatted for chart visualization libraries
func Chart(c fiber.Ctx) error {
	var sales []Sales

	// Execute raw SQL to aggregate daily sales
	// Groups by date and sums (price * quantity) for all order items
	database.DB.Raw(`
		SELECT DATE_FORMAT(o.create_at, '%Y-%m-%d') as date, SUM(oi.price*oi.quantity) as sum 
		FROM orders o 
		JOIN order_items oi on o.id=oi.order_id 
		GROUP BY date
		`).Scan(&sales)
	return c.JSON(sales)
}
