Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Программа выведет: [77 78 79]
Так как в слайс b мы передает срез с индексами от 1 до 4(не включительно,соответсвенно 1;2;3) 

```
