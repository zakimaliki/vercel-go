package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/subosito/gotenv"
	. "github.com/tbxark/g4vercel"
)

var DB *gorm.DB

func InitDB() {
	url := os.Getenv("URL")
	var err error
	DB, err = gorm.Open("postgres", url)
	if err != nil {
		panic("failed to connect database")
	}
}

type Product struct {
	gorm.Model
	Name  string
	Price int
	Stock int
}

func SelectAll() *gorm.DB {
	items := []Product{}
	return DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Product
	return DB.First(&item, "id = ?", id)
}

func Migration() {
	DB.AutoMigrate(&Product{})
}

func ListProducts(context *Context) {
	products := SelectAll()
	context.JSON(200, H{
		"data": products,
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	gotenv.Load()
	InitDB()
	Migration()
	defer DB.Close()
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

	server.GET("/products", ListProducts)
	server.Handle(w, r)
}
