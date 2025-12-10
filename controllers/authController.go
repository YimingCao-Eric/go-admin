package controllers

import (
	"golangProject/database"
	"golangProject/models"
	"golangProject/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

// Register handles user registration
// Validates password confirmation, creates a new user account with default role,
// and returns the created user data (password excluded in response)
func Register(c fiber.Ctx) error {
	var data map[string]string

	// Parse JSON request body
	if err := c.Bind().Body(&data); err != nil {
		return err
	}

	// Validate password confirmation matches
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Passwords do not match",
		})
	}

	// Create user instance with provided data
	// RoleId 3 is assigned as the default role for new registrations
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    3,
	}

	// Hash password before storing (uses bcrypt internally)
	user.SetPassword(data["password"])

	// Persist user to database
	database.DB.Create(&user)

	return c.JSON(user)
}

// Login authenticates a user and establishes a session
// Validates email and password, then issues a JWT token stored in an HTTP-only cookie
// Returns success message on successful authentication
func Login(c fiber.Ctx) error {
	var data map[string]string

	// Parse JSON request body
	if err := c.Bind().Body(&data); err != nil {
		return err
	}

	var user models.User

	// Look up user by email address
	database.DB.Where("email = ?", data["email"]).First(&user)

	// Verify user exists (Id == 0 indicates no record found)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "email not found",
		})
	}

	// Verify password against stored hash (uses bcrypt internally)
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "incorrect password",
		})
	}

	// Generate JWT token containing user ID
	token, err := util.GenerateJWT(strconv.Itoa(int(user.Id)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Set JWT token in HTTP-only cookie (24 hour expiration)
	// HTTP-only prevents client-side JavaScript access for security
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success login",
	})
}

// User retrieves the current authenticated user's profile
// Extracts user ID from JWT token in cookie and returns user data
// Password field is automatically excluded from response via JSON tag
func User(c fiber.Ctx) error {
	// Extract JWT token from authentication cookie
	cookie := c.Cookies("jwt")

	// Parse token and extract user ID from claims
	id, _ := util.ParseJWT(cookie)

	var user models.User

	// Retrieve user record by ID
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

// Logout invalidates the user session by clearing the JWT cookie
// Sets cookie expiration to past time and empty value to force browser deletion
func Logout(c fiber.Ctx) error {
	// Clear JWT cookie by setting empty value and past expiration
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success logout",
	})
}

// UpdateInfo updates the authenticated user's personal information
// Allows users to modify their first name, last name, and email
// User ID is extracted from JWT token to ensure users can only update their own data
func UpdateInfo(c fiber.Ctx) error {
	var data map[string]string

	// Parse JSON request body
	if err := c.Bind().Body(&data); err != nil {
		return err
	}

	// Extract user ID from JWT token in authentication cookie
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJWT(cookie)
	userId, _ := strconv.Atoi(id)

	// Prepare user instance with ID and updated fields
	user := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	// Update user record in database
	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

// UpdatePassword changes the authenticated user's password
// Requires password confirmation to prevent typos
// User ID is extracted from JWT token to ensure users can only change their own password
func UpdatePassword(c fiber.Ctx) error {
	var data map[string]string

	// Parse JSON request body
	if err := c.Bind().Body(&data); err != nil {
		return err
	}

	// Validate password confirmation matches
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Passwords do not match",
		})
	}

	// Extract user ID from JWT token in authentication cookie
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJWT(cookie)
	userId, _ := strconv.Atoi(id)

	// Create user instance with ID only for targeted update
	user := models.User{
		Id: uint(userId),
	}

	// Hash new password before storing (uses bcrypt internally)
	user.SetPassword(data["password"])

	// Update password field in database
	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}
