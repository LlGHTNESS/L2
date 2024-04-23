Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Поскольку test() возвращает nil, который имеет тип *customError, при присваивании переменной err интерфейсного типа, интерфейс err хранит nil в качестве значения и *customError в качестве типа. Таким образом, интерфейс err не равен nil. Поэтому проверка if err != nil возвращает true, и программа выводит "error".


```
