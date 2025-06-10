package main

import "fmt"

type Product interface {
	Price() float64
}

const (
	SmallProduct  = "Small"
	MediumProduct = "Medium"
	LargeProduct  = "Large"
)

type small struct {
	price float64
}

type medium struct {
	price float64
}

type large struct {
	price float64
}

func (s small) Price() float64 {
	return s.price
}

func (m medium) Price() float64 {
	return m.price
}

func (l large) Price() float64 {
	return l.price
}

func createProduct(productType string, price float64) Product {
	switch productType {
	case SmallProduct:
		return small{price: price}
	case MediumProduct:
		return medium{price: price + (price * .03) + (price * .03)}
	case LargeProduct:
		return large{price + (price * .06) + 2500}
	}
	return nil
}

func main() {
	smallProduct := createProduct(SmallProduct, 1000)
	mediumProduct := createProduct(MediumProduct, 1000)
	largeProduct := createProduct(LargeProduct, 1000)

	fmt.Println("Small Product Price:", smallProduct.Price())
	fmt.Println("Medium Product Price:", mediumProduct.Price())
	fmt.Println("Large Product Price:", largeProduct.Price())
}
