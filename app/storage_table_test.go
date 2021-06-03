package app_test

import (
	"math/rand"
	"testing"

	"github.com/joubertredrat/pokelang/app"
)

func TestStorageTable(t *testing.T) {
	tests := []struct {
		name          string
		storage       app.Storage
		countExpected int
	}{
		{
			name:          "Test storage with 10 items",
			storage:       app.NewMemoryStorage(),
			countExpected: 10,
		},
		{
			name:          "Test storage with 50 items",
			storage:       app.NewMemoryStorage(),
			countExpected: 50,
		},
		{
			name:          "Test storage with 100 items",
			storage:       app.NewMemoryStorage(),
			countExpected: 100,
		},
		{
			name:          "Test storage with random items",
			storage:       app.NewMemoryStorage(),
			countExpected: rand.Intn(499-10) + 499,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := test.storage

			if 0 != storage.Count() {
				t.Errorf("storage.Count() expected 0, got %d", storage.Count())
			}

			for i := 1; i <= test.countExpected; i++ {
				storage.AddOne()
			}

			if uint(test.countExpected) != storage.Count() {
				t.Errorf("storage.Count() expected %d, got %d", test.countExpected, storage.Count())
			}
		})
	}
}
