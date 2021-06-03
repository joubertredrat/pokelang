package app_test

import (
	"math/rand"
	"testing"

	"github.com/joubertredrat/pokelang/app"
)

func TestStorageWith10Items(t *testing.T) {
	storage := app.NewMemoryStorage()

	if 0 != storage.Count() {
		t.Errorf("storage.Count() expected 0, got %d", storage.Count())
	}

	countExpected := 10

	for i := 1; i <= countExpected; i++ {
		storage.AddOne()
	}

	if uint(countExpected) != storage.Count() {
		t.Errorf("storage.Count() expected %d, got %d", countExpected, storage.Count())
	}
}

func TestStorageWith50Items(t *testing.T) {
	storage := app.NewMemoryStorage()

	if 0 != storage.Count() {
		t.Errorf("storage.Count() expected 0, got %d", storage.Count())
	}

	countExpected := 50

	for i := 1; i <= countExpected; i++ {
		storage.AddOne()
	}

	if uint(countExpected) != storage.Count() {
		t.Errorf("storage.Count() expected %d, got %d", countExpected, storage.Count())
	}
}

func TestStorageWith100Items(t *testing.T) {
	storage := app.NewMemoryStorage()

	if 0 != storage.Count() {
		t.Errorf("storage.Count() expected 0, got %d", storage.Count())
	}

	countExpected := 100

	for i := 1; i <= countExpected; i++ {
		storage.AddOne()
	}

	if uint(countExpected) != storage.Count() {
		t.Errorf("storage.Count() expected %d, got %d", countExpected, storage.Count())
	}
}

func TestStorageWithRandomItems(t *testing.T) {
	storage := app.NewMemoryStorage()

	if 0 != storage.Count() {
		t.Errorf("storage.Count() expected 0, got %d", storage.Count())
	}

	countExpected := rand.Intn(499-10) + 499

	for i := 1; i <= countExpected; i++ {
		storage.AddOne()
	}

	if uint(countExpected) != storage.Count() {
		t.Errorf("storage.Count() expected %d, got %d", countExpected, storage.Count())
	}
}
