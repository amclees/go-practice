package heap

import "testing"

func TestMinHeap(t *testing.T) {
  m := MinHeap{}

  m.Add(int64(5))
  m.Add(int64(3))
  m.Add(int64(10))
  m.Add(int64(7))
  m.Add(int64(6))

  s := make([]int64, 5)
  for i := range s {
    s[i] = m.Extract()
  }

  b := []int64{3, 5, 6, 7, 10}
  for i := range s {
    if s[i] != b[i] {
      t.Errorf("Expected s[i] to be %d, got %d", b[i], s[i])
    }
  }
}
