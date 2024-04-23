package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Интерфейс состояния
type State interface {
	Handle(context *Context)
}

// Конкретная реализация состояния #1
type ConcreteStateA struct{}

func (csa *ConcreteStateA) Handle(context *Context) {
	fmt.Print("Print")
	context.SetState(&ConcreteStateB{})
}

// Конкретная реализация состояния #2
type ConcreteStateB struct{}

func (csb *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Print2")
	context.SetState(&ConcreteStateA{})
}

// Контекст
type Context struct {
	state State
}

func NewContext(state State) *Context {
	return &Context{state}
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

/* Клиентский код
func main() {
	// Создание объектов состояний
	stateA := &ConcreteStateA{}

	// Создание объекта контекста с начальным состоянием
	context := NewContext(stateA)

	// Использование контекста для выполнения действий в зависимости от состояния
	context.Request()
	context.Request()
	context.Request()
}
Применимость паттерна:
Когда поведение объекта зависит от его состояния и должно быть возможно изменить его поведение во время выполнения, в ответ на изменение состояния.
Когда операции содержат большие условные инструкции, которые зависят от состояния объекта – паттерн помогает избавиться от этих условий.
Плюсы паттерна "Состояние":
Позволяет объекту изменять свое поведение, когда изменяется его внутреннее состояние, кажется, что объект меняет свой класс.
Изоляция кода для каждого состояния и переходов между состояниями в отдельные классы.
Упрощение кода, путем избавления от многочисленных условных операторов.
Минусы паттерна "Состояние":
Может быть излишним, если состояний мало и редко меняются.
Необходимо постоянно поддерживать соответствие между состояниями и переходами, что может усложнить код.
Повышение сложности кода за счет дополнительного количества классов.
*/
