package sort

import (
	"math/rand"
	"testing"
)

func TestShellsort(t *testing.T) {
	a := []int64{6, 2, 5, 2, 4, 3}
	b := []int64{2, 2, 3, 4, 5, 6}
	Shellsort(a)
	fail := false
	for i, _ := range a {
		if a[i] != b[i] {
			fail = true
			break
		}
	}
	if fail {
		t.Errorf("Expected a to be %v, got %v", b, a)
	}
}

func BenchmarkShellsort(b *testing.B) {
	s := make([]int64, b.N)
	for i := range s {
		s[i] = rand.Int63()
	}

	b.ResetTimer()

	Shellsort(s)
}
