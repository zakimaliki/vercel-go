package handler

import (
	"fmt"
	"net/http"

	. "github.com/tbxark/g4vercel"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var products = []Product{
	{Name: "Product A", Price: 100000, Stock: 100},
	{Name: "Product B", Price: 140000, Stock: 50},
	{Name: "Product C", Price: 123000, Stock: 200},
	// Add more sample products as needed
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "OK",
		})
	})

	server.GET("/hello", func(context *Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})

	server.GET("/products", func(context *Context) {
		context.JSON(200, H{
			"data": products,
		})
	})

	server.GET("/product/:id", func(context *Context) {
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
	})

	server.Handle(w, r)
}
