package controllers

import (
	. "github.com/tbxark/g4vercel"
	"github.com/tbxark/g4vercel-demo/api/src/models"
)

func ListProducts(context *Context) {
	products := models.SelectAll()
	context.JSON(200, H{
		"data": products,
	})
}

// func DetailProduct(context *Context) {
// 	id := context.Param("id")
// 	index := -1
// 	for i, product := range products {
// 		if product.Name == id {
// 			index = i
// 			break
// 		}
// 	}
// 	if index == -1 {
// 		context.JSON(404, H{
// 			"message": "Product not found",
// 		})
// 	} else {
// 		context.JSON(200, H{
// 			"data": products[index],
// 		})
// 	}
// }
