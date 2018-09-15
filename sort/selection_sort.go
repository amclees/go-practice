package sort

import "math"

func SelectionSort(s []int64) {
  for i := range s {
    var m int64 = math.MaxInt64
    w := -1
    for j := i + 1; j < len(s); j++ {
      if s[j] < m {
        m = s[j]
        w = j
      }
    }
    if w == -1 {
      return
    }
    k := s[i]
    s[i] = m
    s[w] = k
  }
}
