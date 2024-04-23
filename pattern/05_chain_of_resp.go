package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

// Базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler
}

func (bh *BaseHandler) SetNext(handler Handler) Handler {
	bh.nextHandler = handler
	return handler
}

func (bh *BaseHandler) Handle(request string) {
	if bh.nextHandler != nil {
		bh.nextHandler.Handle(request)
	}
}

// Конкретные реализации обработчиков
type FirstHandler struct {
	BaseHandler
}

func (fh *FirstHandler) Handle(request string) {
	if request == "first" {
		fmt.Println("FirstHandler handled the request")
		return
	}
	fh.BaseHandler.Handle(request)
}

type SecondHandler struct {
	BaseHandler
}

func (sh *SecondHandler) Handle(request string) {
	if request == "second" {
		fmt.Println("SecondHandler handled the request")
		return
	}
	sh.BaseHandler.Handle(request)
}

type ThirdHandler struct {
	BaseHandler
}

func (th *ThirdHandler) Handle(request string) {
	if request == "third" {
		fmt.Println("ThirdHandler handled the request")
		return
	}
	th.BaseHandler.Handle(request)
}

/* Клиентский код
func main() {
	// Создание объектов обработчиков
	firstHandler := &FirstHandler{}
	secondHandler := &SecondHandler{}
	thirdHandler := &ThirdHandler{}

	// Установка цепочки вызовов
	firstHandler.SetNext(secondHandler).SetNext(thirdHandler)

	// Выполнение запросов
	firstHandler.Handle("second")
	firstHandler.Handle("third")
	firstHandler.Handle("fourth")
}
Применимость паттерна:
Когда имеется более одного объекта, который может обработать определённый запрос, и конкретный обработчик заранее неизвестен.
Когда нужно отправить запрос нескольким объектам, без указания конкретных получателей.
Когда набор обработчиков может формироваться динамически.
Плюсы паттерна "Цепочка обязанностей":
Уменьшает зависимость между клиентом и обработчиками, позволяя добавлять новые обработчики в любой момент.
Реализует принцип единственной обязанности, разделяя разные обязанности на разные классы.
Обеспечивает гибкость при распределении обязанностей между объектами.
Минусы паттерна "Цепочка обязанностей":
Запрос может закончиться без обработки, если в цепочке нет подходящего обработчика.
Порой сложно проследить путь запроса, особенно в сложных системах, что может усложнить отладку.
Может повлечь за собой незамеченные затраты производительности, если цепочка длинная или обработка запроса тяжёлая.
*/
