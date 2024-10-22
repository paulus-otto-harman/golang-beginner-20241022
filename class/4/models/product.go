package models

type Product struct {
	name     string
	quantity int
}

func InitProduct(name string, quantity int) Product {
	return Product{name, quantity}
}

func (product *Product) Nama() string {
	return product.name
}

func (product *Product) Quantity() int {
	return product.quantity
}

func (product *Product) SetQuantity(kuantiti int) {
	product.quantity = kuantiti
}
