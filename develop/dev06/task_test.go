package main

import (
	"os"
	"reflect"
	"testing"
)

// Тест функции parseFields
func TestParseFields(t *testing.T) {
	testCases := []struct {
		input  string
		expect map[int]bool
	}{
		{"1,2,3", map[int]bool{1: true, 2: true, 3: true}},
		{"4,5", map[int]bool{4: true, 5: true}},
	}

	for i, tc := range testCases {
		result, err := parseFields(tc.input)
		if err != nil {
			t.Errorf("Case %d: unexpected error: %s", i, err)
		}
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("Case %d: expected %v, got %v", i, tc.expect, result)
		}
	}

	// Проверка некорректного ввода
	_, err := parseFields("a,b")
	if err == nil {
		t.Error("Expected error for non-integer input, got nil")
	}
}

// Пример мокирования STDIN и тестирования основной логики (деструктурирование строки на колонки и выбор указанных полей)
// Для выполнения этого теста требуется дополнительная реализация или изменение существующей функции для тестирования
func TestCutColumns(t *testing.T) {
	input := "a\tb\tc\nd\te\tf"
	oldStdin := os.Stdin                   // Сохраняем стандартный ввод для последующего восстановления
	defer func() { os.Stdin = oldStdin }() // Восстанавливаем стандартный ввод после теста

	// Создаем временный файл для мокирования STDIN
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %s", err)
	}
	defer os.Remove(tmpfile.Name()) // очищаем после себя

	if _, err := tmpfile.WriteString(input); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatalf("Failed to seek to start of file: %s", err)
	}

	os.Stdin = tmpfile

	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %s", err)
	}
}
