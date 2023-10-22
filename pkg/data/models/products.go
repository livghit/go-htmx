package models

type Product struct {
	Id          int32
	Name        string
	Description string
}

// Easy Way of creating a Product
func CreateProduct(name string, description string) Product {
	return Product{Name: name, Description: description}
}
