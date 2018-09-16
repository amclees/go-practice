package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuicksort(t *testing.T) {
	a := []int64{6, 2, 5, 2, 4, 3}
	b := []int64{2, 2, 3, 4, 5, 6}
	Quicksort(a)
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

func BenchmarkQuicksort(b *testing.B) {
	s := make([]int64, b.N)
	for i := range s {
		s[i] = rand.Int63()
	}

	b.ResetTimer()

	Quicksort(s)
}

type Int64Slice []int64

func (s Int64Slice) Len() int {
	return len(s)
}

func (s Int64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Int64Slice) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

func BenchmarkBuiltinSort(b *testing.B) {
	s := make([]int64, b.N)
	for i := range s {
		s[i] = rand.Int63()
	}

	b.ResetTimer()

	sort.Sort(Int64Slice(s))
}
