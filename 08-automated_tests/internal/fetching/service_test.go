package fetching

import (
	"testing"

	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/storage/inmem"
)

func TestFetchByID(t *testing.T) {
	repo := inmem.NewRepository()

	service := NewService(repo)

	expected := 127
	b, err := service.FetchByID(expected)
	if err != nil {
		t.Fatalf("expected %d, got an error %v", expected, err)
	}

	if b.ProductID != expected {
		t.Fatalf("expected %d,  got: %d ", expected, b.BeerID)
	}
}
