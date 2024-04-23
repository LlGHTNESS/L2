Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
каналы, используемые в функции merge, не закрываются после завершения передачи значений. Это приводит к тому, что после передачи всех значений из исходных каналов, функция merge продолжит выполнение и будет возвращать нули, поскольку попытка чтения из закрытого канала сразу возвращает нулевое значение для типа канала. Таким образом, в главной горутине цикл for v := range c никогда не завершится, поскольку канал c остается открытым и продолжает отправлять нулевые значения.

```
