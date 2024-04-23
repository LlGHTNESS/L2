package main

import (
	"os"
	"reflect"
	"testing"
)

// Тестирование функции ReadFileStrings
func TestReadFileStrings(t *testing.T) {
	// Создаем временный тестовый файл
	const testFileName = "test.txt"
	const fileContent = "Строка1\nСтрока2\nСтрока3"
	err := os.WriteFile(testFileName, []byte(fileContent), 0666)
	if err != nil {
		t.Fatal("Не удалось создать тестовый файл:", err)
	}
	defer os.Remove(testFileName)

	// Чтение содержимого файла с помощью ReadFileStrings
	strings, err := ReadFileStrings(testFileName)
	if err != nil {
		t.Fatal("Ошибка во время чтения файла :", err)
	}

	// Ожидаемое содержимое файла
	expectedStrings := []string{"Строка1", "Строка2", "Строка3"}

	// Проверяем, соответствует ли фактическое содержимое файлов ожидаемому
	if !reflect.DeepEqual(strings, expectedStrings) {
		t.Errorf("ReadFileStrings = %q, ожидалось %q", strings, expectedStrings)
	}
}

// Тест парсинга аргументов:
func TestParseArguments(t *testing.T) {
	// Сохраняем исходное значение os.Args
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Тестовые аргументы
	os.Args = []string{"cmd", "-k", "2", "-n", "-r", "-u"}

	// Создаем экземпляр Params и парсим аргумены
	params := Params{}
	params.ParseArguments()

	// Ожидаемые значения Params
	expectedParams := Params{
		SortColumn:       2,
		SortByNums:       true,
		SortReverse:      true,
		RemoveDuplicates: true,
	}

	// Проверка значений
	if params != expectedParams {
		t.Errorf("ParseArguments() = %#v, ожидалось %#v", params, expectedParams)
	}
}
