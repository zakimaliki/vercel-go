package routes

import (
	"fmt"

	. "github.com/tbxark/g4vercel"
	"github.com/tbxark/g4vercel-demo/api/src/controllers"
)

func Router(server *Engine) {
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

	server.GET("/products", controllers.ListProducts)

	// server.GET("/product/:id", controllers.DetailProduct)

}
