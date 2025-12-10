package controllers

import (
	"github.com/gofiber/fiber/v3"
)

// Upload handles file uploads via multipart form data
// Accepts files from clients, saves them to the server's uploads directory,
// and returns a publicly accessible URL for the uploaded file
// Expected form field name: "image"
func Upload(c fiber.Ctx) error {
	// Parse multipart form data from request body
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Extract files from the "image" form field
	// Supports multiple file uploads in a single request
	files := form.File["image"]

	fileName := ""

	// Process each uploaded file
	for _, file := range files {
		fileName = file.Filename

		// Save file to ./uploads/ directory with original filename
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return err
		}
	}

	// Return public URL for accessing the uploaded file
	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + fileName,
	})
}
