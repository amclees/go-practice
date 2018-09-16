package sort

func Heapsort(s []int64) {
  makeHeap(s)

  ss := s[:]
  for ; len(ss) > 0; {
    ss = ss[1:]

    minHeapify(ss, 0)
  }
}

func makeHeap(s []int64) {
  for i := range s {
    for {
      ni := i / 2
      if s[ni] > s[i] {
        a := s[ni]
        s[ni] = s[i]
        s[i] = a
      }

      if i == 0 {
        break
      }
      i = ni
    }
  }
}

func minHeapify(s []int64, i int) {
  l := (i * 2) + 1
  r := l + 1

  swap := -1
  if l < len(s) && s[l] < s[i] {
    swap = l
  }
  if r < len(s) && s[r] < s[i] && (swap == -1 || s[swap] > s[r]) {
    swap = r
  }

  if swap != -1 {
    a := s[swap]
    s[swap] = s[i]
    s[i] = a

    minHeapify(s, swap)
  }
}
