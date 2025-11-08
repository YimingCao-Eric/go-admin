package controllers

import (
	"github.com/gofiber/fiber/v3"
)

// Upload handles file uploads via multipart form data
// This function receives files from clients, saves them to the server's uploads directory,
// and returns a URL where the uploaded file can be accessed
func Upload(c fiber.Ctx) error {
	// Parse the multipart form data from the request
	// Multipart forms are used when uploading files via HTML forms or API clients
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Extract all files from the "image" form field
	// This expects the client to send files with the field name "image"
	// The form.File map returns a slice of uploaded files for the given field name
	files := form.File["image"]

	// Variable to store the saved filename
	fileName := ""

	// Process each uploaded file in the "image" field
	// This loop handles multiple file uploads if the client sends multiple files
	for _, file := range files {
		// Store the original filename from the uploaded file
		fileName = file.Filename

		// Save the file to the server's uploads directory
		// The file is saved with its original filename in the ./uploads/ folder
		// c.SaveFile handles the actual file copying from memory to disk
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return err
		}
	}

	// Return a JSON response containing the public URL where the file can be accessed
	// Clients can use this URL to retrieve the uploaded file later
	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + fileName,
	})
}
