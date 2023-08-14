package models

import (
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Date string `json:"date"`
	Link string `json:"link"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}