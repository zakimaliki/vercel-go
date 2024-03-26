package models

import (
	"github.com/jinzhu/gorm"
	"github.com/tbxark/g4vercel-demo/api/src/config"
)

type Product struct {
	gorm.Model
	Name  string
	Price int
	Stock int
}

func SelectAll() *gorm.DB {
	items := []Product{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Product
	return config.DB.First(&item, "id = ?", id)
}
