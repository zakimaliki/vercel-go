package helper

import (
	"github.com/tbxark/g4vercel-demo/api/src/config"
	"github.com/tbxark/g4vercel-demo/api/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Product{})
}
