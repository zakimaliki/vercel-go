package controllers

import (
	. "github.com/tbxark/g4vercel"
	"github.com/tbxark/g4vercel-demo/api/src/models"
)

var products = []models.Product{
	{Name: "Product A", Price: 100000, Stock: 100},
	{Name: "Product B", Price: 140000, Stock: 50},
	{Name: "Product C", Price: 123000, Stock: 200},
	// Add more sample products as needed
}

func ListProducts(context *Context) {
	context.JSON(200, H{
		"data": products,
	})
}

func DetailProduct(context *Context) {
	id := context.Param("id")
	index := -1
	for i, product := range products {
		if product.Name == id {
			index = i
			break
		}
	}
	if index == -1 {
		context.JSON(404, H{
			"message": "Product not found",
		})
	} else {
		context.JSON(200, H{
			"data": products[index],
		})
	}
}
