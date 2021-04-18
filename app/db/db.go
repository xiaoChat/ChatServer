package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("init start ...")
}

func GetYexk() (db *gorm.DB, err error) {
	dsn := "root:123456@tcp(mysql:3306)/yexk?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
