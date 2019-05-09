package fetching

import (
	"errors"
	"testing"

	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/storage/inmem"
)

func TestFetchByID(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
		err   error
	}{
		"valid beer":     {input: 127, want: 127, err: nil},
		"not found beer": {input: 99999, err: errors.New("error")},
	}
	repo := inmem.NewRepository()
	service := NewService(repo)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := service.FetchByID(tc.input)
			if err != nil && tc.err == nil {
				t.Fatalf("not expected any errors and got %v", err)
			}

			if err == nil && tc.err != nil {
				t.Error("expected an error and got nil")
			}

			if b.ProductID != tc.want {
				t.Fatalf("expected %d,  got: %d ", tc.want, b.ProductID)
			}
		})
	}
}
