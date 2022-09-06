package database

import (	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:password_root@/godb"), &gorm.Config{})

	if err != nil {
		panic("Can't connect.")
	}
}