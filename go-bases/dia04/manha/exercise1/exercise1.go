package main

import "fmt"

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []Product{
	{Id: 1, Name: "Product 1", Price: 10.0, Description: "Description of Product 1", Category: "Category A"},
	{Id: 2, Name: "Product 2", Price: 20.0, Description: "Description of Product 2", Category: "Category B"},
}

func (p Product) GetAll() {
	for _, p := range products {
		fmt.Println("Id:", p.Id, "Name:", p.Name, "Price:", p.Price, "Description:", p.Description, "Category:", p.Category)
	}
}

func (p Product) Save(product Product) {
	products = append(products, product)
}

func (p Product) getById(id int) Product {
	for _, product := range products {
		if product.Id == id {
			return product
		}
	}
	return Product{}
}

func main() {
	product := Product{}
	product.GetAll()
	newProduct := Product{Id: 3, Name: "Product 3", Price: 30.0, Description: "Description of Product 3", Category: "Category C"}
	product.Save(newProduct)
	product.GetAll()
	product2 := product.getById(2)
	fmt.Println("Product found:", product2)
}
