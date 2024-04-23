package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func printLine(line string, lineNum int, prtintLineNum bool) {
	if prtintLineNum {
		fmt.Println(strconv.Itoa(lineNum) + ":" + line)
	} else {
		fmt.Println(line)
	}
}

func main() {
	// Парсим параметры
	params := Params{}
	err := params.ParseArguments()
	if err != nil {
		log.Fatal("Can't parce arguments: " + string(err.Error()))
	}

	if flag.NArg() < 2 {
		log.Fatal("No pattern or file name found")
	}

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	// Открываем файл для чтения
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Can't open file: " + string(err.Error()))
	}
	defer file.Close()

	//Пременные для вывода информации
	beforeLines := NewCircularQueue[string](params.Before) // Очередь предыдущих строк
	currentLineNum := 0
	afterLinesCount := 0
	matchedCount := 0

	// Компиляция регулярного выражения
	var re *regexp.Regexp
	if params.Fixed { // Если необходимо точное значение, преобразуем шаблон
		pattern = regexp.QuoteMeta(pattern)
	}
	if params.IgnoreCase {
		re, err = regexp.Compile("(?i)" + pattern) // Игнорирование регистра
	} else {
		re, err = regexp.Compile(pattern) // Обычное выражение
	}
	if err != nil {
		log.Fatal("Can't compile regexp: " + string(err.Error()))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()          // Читаем строку
		matched := re.MatchString(currentLine) // Есть ли совпадение в строке?

		// Если строка необходимая
		if (matched && !params.Invert) || (!matched && params.Invert) {
			// Если считаем строки
			if params.Counting {
				matchedCount += 1
				continue
			}

			// Выводим строки до
			for !beforeLines.IsEmpty() {
				printLine(beforeLines.Pull(), currentLineNum-beforeLines.Len()-1, params.LineNum)
			}

			// Выводим данную строку
			printLine(currentLine, currentLineNum, params.LineNum)

			// Сообщаем что нужно писать контекст
			afterLinesCount = params.After
		} else {
			if afterLinesCount <= 0 { // Сохраняем стороку на случай нахождения
				beforeLines.Push(currentLine)
			} else { // Или печатаем контекст если необходимо
				afterLinesCount -= 1
				printLine(currentLine, currentLineNum, params.LineNum)
			}
		}
		currentLineNum += 1
	}

	// Вывод количества строк
	if params.Counting {
		fmt.Println(matchedCount)
	}
}

type Params struct {
	After  int
	Before int
	// Context    int
	Counting   bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func (p *Params) ParseArguments() error {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать +-N строк вокруг совпадения")
	count := flag.Bool("c", false, "подсчитать количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	if *context != 0 {
		if *after != 0 || *before != 0 {
			return errors.New("-A and -B mutually exclusive with -C")
		} else {
			*after = *context
			*before = *context
		}
	}

	p.After = *after
	p.Before = *before
	// p.Context = *context
	p.Counting = *count
	p.IgnoreCase = *ignoreCase
	p.Invert = *invert
	p.Fixed = *fixed
	p.LineNum = *lineNum

	return nil
}

type CircularQueue[T any] struct {
	buffer []T
	head   int
	tail   int
	size   int
	count  int
}

func NewCircularQueue[T any](size int) *CircularQueue[T] {
	return &CircularQueue[T]{
		buffer: make([]T, size),
		size:   size,
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (q *CircularQueue[T]) Push(item T) {
	if q.size == 0 {
		return
	}
	if q.count == q.size {
		q.head = (q.head + 1) % q.size
	} else {
		q.count++
	}
	q.buffer[q.tail] = item
	q.tail = (q.tail + 1) % q.size

}

func (q *CircularQueue[T]) Pull() T {
	if q.count == 0 {
		var zero T
		return zero
	}
	q.count--
	item := q.buffer[q.head]
	q.head = (q.head + 1) % q.size
	return item
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *CircularQueue[T]) Clear() {
	q.head = 0
	q.tail = 0
	q.count = 0
}

func (q *CircularQueue[T]) Len() int {
	return q.count
}
