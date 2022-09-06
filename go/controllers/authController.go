package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lgustavopalmieri/go-admin-api/go/database"
	"github.com/lgustavopalmieri/go-admin-api/go/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	// here is where we get the payload data
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords doesn't match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	// here is where we get the payload data
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//

	var user models.User

	// searchs for an user in database (like typeorm where)
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password or email aren't corrects. Try again.",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password or email aren't corrects. Try again.",
		})
	}

	return c.JSON(user)
}
