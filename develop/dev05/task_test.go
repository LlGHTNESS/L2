package main

import (
	"testing"
)

func TestCircularQueuePushPull(t *testing.T) {
	q := NewCircularQueue[string](3)

	// Тестируем Push и Pull на пустой очереди
	if !q.IsEmpty() {
		t.Errorf("New queue should be empty but got %v", q.IsEmpty())
	}

	// Тест данных
	testData := []string{"one", "two", "three", "four"}

	// Пушим данные в очередь и проверяем результат метода Pull
	for _, v := range testData {
		q.Push(v)
		if q.Pull() != v {
			t.Errorf("Expected %v but got %v", v, q.Pull())
		}
	}

	// Проверяем пустоту очереди после всех операций
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty but got %v", q.Len())
	}
}

func TestCircularQueueOverflow(t *testing.T) {
	q := NewCircularQueue[string](2)

	testData := []string{"one", "two", "three"}

	for _, v := range testData {
		q.Push(v)
	}

	if q.Len() > 2 {
		t.Errorf("Queue length should not exceed 2 but got %v", q.Len())
	}
}
