package main

import "fmt"

// Интерфейс фабрики
type Factory interface {
	create() Product
}

// Интерфейс продукта
type Product interface {
	use()
}

// Фабрика для создания продукта "Продукт А"
type FactoryA struct{}

func (f *FactoryA) create() Product {
	return &ProductA{}
}

// Продукт "Продукт А"
type ProductA struct{}

func (p *ProductA) use() {
	fmt.Println("Используется продукт А")
}

// Фабрика для создания продукта "Продукт В"
type FactoryB struct{}

func (f *FactoryB) create() Product {
	return &ProductB{}
}

// Продукт "Продукт В"
type ProductB struct{}

func (p *ProductB) use() {
	fmt.Println("Используется продукт В")
}

func main() {
	factoryA := &FactoryA{}
	factoryB := &FactoryB{}

	productA := factoryA.create()
	productB := factoryB.create()

	productA.use() // Используется продукт А
	productB.use() // Используется продукт В
}
