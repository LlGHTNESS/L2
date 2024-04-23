package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Интерфейс посетителя
type Visitor interface {
	VisitHubaBuba(element *HubaBuba)
	VisitTakSak(element *TakSak)
}

// Конкретный посетитель 1
type JSONEncoder struct{}

func (cv1 *JSONEncoder) VisitHubaBuba(element *HubaBuba) {
	fmt.Println("JSON HubaBuba")
}

func (cv1 *JSONEncoder) VisitTakSak(element *TakSak) {
	fmt.Println("JSON TakSak")
}

// Конкретный посетитель 2
type XMLEncoder struct{}

func (cv2 *XMLEncoder) VisitHubaBuba(element *HubaBuba) {
	fmt.Println("XML HubaBuba")
}

func (cv2 *XMLEncoder) VisitTakSak(element *TakSak) {
	fmt.Println("XML TakSak")
}

// Интерфейс элемента
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type HubaBuba struct{}

func (ea *HubaBuba) Accept(visitor Visitor) {
	visitor.VisitHubaBuba(ea)
}

// Конкретный элемент B
type TakSak struct{}

func (eb *TakSak) Accept(visitor Visitor) {
	visitor.VisitTakSak(eb)
}

/* Клиентский код
func main() {
	// Создание объектов элементов
	hubaBuba := &HubaBuba{}
	elementB := &TakSak{}

	// Создание объектов посетителей
	jsonEncoder := &JSONEncoder{}
	xmlEncoder := &XMLEncoder{}

	// Применение посетителя к элементам
	hubaBuba.Accept(jsonEncoder)
	elementB.Accept(xmlEncoder)

	hubaBuba.Accept(jsonEncoder)
	elementB.Accept(xmlEncoder)
}
Применимость паттерна "Посетитель":
Когда у вас есть сложная структура объектов (например, дерево), и вы хотите выполнить операции на них, не загрязняя их классы.
Если необходимо добавить функционал к библиотеке или фреймворку, и вы не можете или не хотите изменять существующие классы.
Когда операции, которые нужно выполнить, специфические или редко изменяются, но объекты, над которыми они выполняются, часто обновляются или расширяются.
Плюсы паттерна "Посетитель":
Позволяет добавить новые операции к комплексным структурам объектов, не изменяя классы этих объектов.
Собирает родственные операции и разбросанные по разным классам.
Упрощает добавление операций, работающих со сложными структурами объектов, позволяя избежать "загрязнения" их кода.
Минусы паттерна "Посетитель":
Нарушает инкапсуляцию классов, поскольку "посетитель" знает о внутренностях разных элементов.
При добавлении новых элементов в структуру нужно обновлять всех "посетителей", что может быть проблематично если есть много разных посетителей.
"Посетители" могут стать слишком сложными, если логика операций громоздкая.
*/
