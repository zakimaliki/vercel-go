package models

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
