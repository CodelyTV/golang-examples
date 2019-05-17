package ontario

import "testing"

func BenchmarkGetBeers(b *testing.B) {
	repo := NewOntarioRepository()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		repo.GetBeers()
	}
}
