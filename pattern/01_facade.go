package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type CodeEditor struct{}

func (t *CodeEditor) Save() {
	fmt.Println("Code saved")
}

type Compiller struct{}

func (t *Compiller) Compile() {
	fmt.Println("Code compiled")
}

type CLI struct{}

func (t *CLI) Run() {
	fmt.Println("Program run")
}

type IDEFacade struct {
	codeEditor *CodeEditor
	compiller  *Compiller
	cli        *CLI
}

func newIDEFacade() *IDEFacade {
	ideFacacde := &IDEFacade{
		codeEditor: &CodeEditor{},
		compiller:  &Compiller{},
		cli:        &CLI{},
	}
	return ideFacacde
}

func (ide *IDEFacade) BuildAndRun() {
	ide.codeEditor.Save()
	ide.compiller.Compile()
	ide.cli.Run()
}

/*
func main() {
	ide := newIDEFacade()
	ide.BuildAndRun()
}

Применимость паттерна "Фасад":
Когда у вас есть сложная система, и вы хотите ограничить её сложность для клиентов.
Когда нужно разделить систему на слои, так как фасады могут действовать как точки взаимодействия между слоями.
Плюсы паттерна "Фасад":
Упрощение интерфейса: клиенты используют простой интерфейс вместо прямого управления сложными подсистемами.
Разделение слоёв: клиенты зависят от слоёв высокого уровня, а не от подсистем низкого уровня.
Уменьшение сложности кода: клиентский код становится более чистым и понятным.
Минусы паттерна "Фасад":
Риск создания "Божественного объекта": Фасад может стать слишком «тяжёлым», беря на себя больше ответственности, чем следовало бы.
Ограниченная гибкость: для тех клиентов, которым нужен доступ к сложной функциональности подсистемы, фасад может представлять проблему, так как предоставляет ограниченный функционал.
*/
