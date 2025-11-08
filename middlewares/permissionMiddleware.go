package middlewares

import (
	"errors"
	"golangProject/database"
	"golangProject/models"
	"golangProject/util"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func IsAuthorized(c fiber.Ctx, page string) error {
	// Extract JWT token from the "jwt" cookie
	cookie := c.Cookies("jwt")

	Id, err := util.ParseJWT(cookie)
	// Validate the token using the utility function
	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(Id)
	user := models.User{
		Id: uint(userId),
	}
	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}
	database.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}
	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
