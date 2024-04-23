package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Интерфейс стратегии
type Strategy interface {
	Execute()
}

// Конкретная реализация стратегии #1
type MagicSortOne struct{}

func (csa *MagicSortOne) Execute() {
	fmt.Println("Magic sort one")
}

// Конкретная реализация стратегии #2
type MagicSortTwo struct{}

func (csb *MagicSortTwo) Execute() {
	fmt.Println("Magic sort two")
}

// Контекст
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

/* Клиентский код
func main() {
	// Создание объектов стратегий
	magicSortA := &MagicSortOne{}
	magicSortB := &MagicSortTwo{}

	// Создание объектов контекста с разными стратегиями
	contextA := NewContext(magicSortA)
	contextB := NewContext(magicSortB)

	// Использование контекста с разными стратегиями
	contextA.ExecuteStrategy()
	contextB.ExecuteStrategy()
}
Применимость паттерна:
Когда вам нужны различные варианты одного и того же алгоритма.
Если поведение класса часто меняется в зависимости от условий, избегая многочисленных условных операторов.
Для поддержки принципа открытости/закрытости, позволяя классам быть открытыми для расширения, но закрытыми для изменений.
Плюсы паттерна "Стратегия":
Обеспечивает возможность динамической замены алгоритмов во время выполнения программы.
Позволяет изолировать код и данные алгоритмов от других частей программы.
Упрощает юнит-тестирование, так как каждую стратегию можно проверять отдельно.
Уменьшает количество кода внутри классов, связанных с различными поведениями.
Минусы паттерна "Стратегия":
Может увеличить сложность кода из-за введения новых классов.
Клиент должен знать о различиях между стратегиями, чтобы сделать правильный выбор.
*/
