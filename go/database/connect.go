package database

import (
	"github.com/lgustavopalmieri/go-admin-api/go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:password_root@tcp(godb)"), &gorm.Config{})

	if err != nil {
		panic("Can't connect.")
	}

	DB = database

	database.AutoMigrate(&models.User{})
}
