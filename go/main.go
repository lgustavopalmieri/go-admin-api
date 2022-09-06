package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 14 mins
	db, err := gorm.Open(mysql.Open("root:password_root@/godb"), &gorm.Config{})

	if err != nil {
		panic("Can't connect.")
	}

	fmt.Println(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
