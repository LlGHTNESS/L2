package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Обрабатываем параметры запуска
	params := Params{}
	params.ParseArguments()

	// Читаем файл
	fileStrings, err := ReadFileStrings(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal("Can't read file")
	}

	// Сортируем
	stringSorter := NewStringSorter(
		fileStrings,
		params.SortColumn,
		params.SortByNums,
	)
	stringSorter.Sort()

	fileStrings = stringSorter.Get(
		params.SortReverse,
		params.RemoveDuplicates,
	)
	for i := 0; i < len(fileStrings); i++ {
		fmt.Println(fileStrings[i])
	}

}

type Params struct {
	SortColumn       int
	SortByNums       bool
	SortReverse      bool
	RemoveDuplicates bool
}

func (p *Params) ParseArguments() {
	k := flag.Int("k", -1, "Указание колонки для сортировки (слова в строке могут выступатьв качестве колонок, по умолчанию разделитель — пробел)")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Сортировать в обратном порядке")

	flag.Parse()

	p.SortColumn = *k
	p.SortByNums = *n
	p.SortReverse = *r
	p.RemoveDuplicates = *u
}

type StringSorter struct {
	strings    []string
	sortColumn int
	sortByNums bool
}

func NewStringSorter(strings []string, sortColumn int, sortByNums bool) *StringSorter {
	return &StringSorter{
		strings:    strings,
		sortColumn: sortColumn,
		sortByNums: sortByNums,
	}
}
func (ss StringSorter) Len() int      { return len(ss.strings) }
func (ss StringSorter) Swap(i, j int) { ss.strings[i], ss.strings[j] = ss.strings[j], ss.strings[i] }
func (ss StringSorter) stringsLess(a, b string) bool {
	if !ss.sortByNums {
		return a < b // Сравниваем обычные строки
	} else { // Сравниваем строки как числа
		a, err := strconv.Atoi(a)
		if err != nil {
			return true
		}
		b, err := strconv.Atoi(b)
		if err != nil {
			return false
		}
		return a < b
	}
}
func (ss StringSorter) Less(i, j int) bool {
	// Выбираем что сравнивать
	var a, b string
	if ss.sortColumn == -1 {
		a = ss.strings[i]
		b = ss.strings[j]
	} else {
		aColumns := strings.Split(ss.strings[i], " ")
		bColumns := strings.Split(ss.strings[j], " ")

		if len(a) <= ss.sortColumn {
			return true
		} else if len(b) <= ss.sortColumn {
			return false
		}

		a = aColumns[ss.sortColumn]
		b = bColumns[ss.sortColumn]
	}
	// Сравниваем
	return ss.stringsLess(a, b)
}
func (ss StringSorter) Sort() {
	sort.Sort(ss)
}

func (ss StringSorter) Get(reversed bool, removeDuplicates bool) []string {
	result := make([]string, 0)

	begin := 0
	end := 0
	step := 0
	if !reversed {
		begin = 0
		end = len(ss.strings)
		step = 1
	} else {
		begin = len(ss.strings) - 1
		end = -1
		step = -1
	}

	if !removeDuplicates {
		for ; begin != end; begin += step {
			result = append(result, ss.strings[begin])
		}
	} else {
		result = append(result, ss.strings[begin])
		begin += step

		for ; begin != end; begin += step {
			if ss.strings[begin-step] != ss.strings[begin] {
				result = append(result, ss.strings[begin])
			}
		}
	}

	return result
}

func ReadFileStrings(fileName string) (result []string, err error) {
	file, err := os.Open(fileName) // Открываем файл

	if err != nil {
		log.Fatal("Can't open file: " + fileName)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Can't close file")
		}
	}(file)

	// Читаем построчно
	result = make([]string, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	return result, sc.Err()
}

func WriteStringsToFile(fileName string, strings []string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Can't create file")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Can't close file")
		}
	}(file)

	for i := 0; i < len(strings); i++ {
		_, err := fmt.Fprintln(file, strings[i])
		if err != nil {
			log.Fatal("Can't write to file")
		}
	}
}
