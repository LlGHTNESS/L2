package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Интерфейс продукта
type Product interface {
	Use()
}

// Конкретная реализация продукта
type ConcreteProduct struct{}

func (cp *ConcreteProduct) Use() {
	fmt.Println("Using ConcreteProduct")
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

// Конкретная реализация фабрики
type ConcreteFactory struct{}

func (cf *ConcreteFactory) CreateProduct() Product {
	return &ConcreteProduct{}
}

/* Клиентский код
func main() {
	// Создание объекта фабрики
	factory := &ConcreteFactory{}

	// Создание продукта с помощью фабричного метода
	product := factory.CreateProduct()

	// Использование продукта
	product.Use()
}
Применимость паттерна:
Когда заранее неизвестно, объекты каких классов нужно создавать.
Когда система должна быть независимой от процесса создания объектов и их конкретных типов.
Когда нужно предоставить библиотеку классов, показывая только их интерфейсы, а не реализацию.
Плюсы паттерна "Фабричный метод":
Позволяет сделать код создания объектов более общим, не привязываясь к конкретным классам.
Облегчает добавление новых продуктов в программу.
Формирует архитектуру, которая поддерживает принцип открытости/закрытости (classes should be open for extension, but closed for modification).
Минусы паттерна "Фабричный метод":
Может привести к созданию большого числа маленьких классов, что усложняет архитектуру.
Так как фабрика скрывает создающиеся продукты (конкретные классы), использование библиотеки может стать более сложным для понимания извне.
*/
