package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "izasoerya:Makanmieayam4kali@tcp(localhost:3306)/MUSIC_LIST"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}