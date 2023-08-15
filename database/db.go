package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := "izasoerya:Makanmieayam4kali@tcp(localhost:3306)/MUSIC_LIST"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error initializing database")
		panic(err)
	}
	DB = db
	return err
}