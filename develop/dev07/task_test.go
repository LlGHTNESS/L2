package main

import (
	"testing"
	"time"
)

// Канал, который закроется через определенное время
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

// Тест для функции reflect_or
func Test_reflect_or(t *testing.T) {
	start := time.Now()

	// Ожидаем сигнал от самого быстрого канала
	<-reflect_or(
		sig(2*time.Minute),
		sig(1*time.Second),
	)

	if time.Since(start) > 2*time.Second {
		t.Error("reflect_or took too long to receive the signal")
	}
}

// Тест для функции select_or
func Test_select_or(t *testing.T) {
	start := time.Now()

	// Ожидаем сигнал от самого быстрого канала
	<-select_or(
		sig(2*time.Minute),
		sig(1*time.Second),
	)

	if time.Since(start) > 2*time.Second {
		t.Error("select_or took too long to receive the signal")
	}
}

// Тест для функции goroutines_or
func Test_goroutines_or(t *testing.T) {
	start := time.Now()

	// Ожидаем сигнал от самого быстрого канала
	<-gorutines_or(
		sig(2*time.Minute),
		sig(1*time.Second),
	)

	if time.Since(start) > 2*time.Second {
		t.Error("goroutines_or took too long to receive the signal")
	}
}
