package sort

func MergeSort(s []int64) {
  if len(s) < 2 {
    return
  }
  m := len(s) / 2
  MergeSort(s[0:m])
  MergeSort(s[m:])
  merge(s, m)
}

func merge(s []int64, m int) {
  b := make([]int64, len(s))
  p1 := m - 1
  p2 := len(s) - 1

  finish := -1
  fi := 0
  for i := range b {
    v1 := s[p1]
    v2 := s[p2]

    if v1 >= v2 {
      p1 -= 1
      b[len(b) - 1 - i] = v1
      if p1 < 0 {
        finish = p2
        fi = i
        break
      }
    } else {
      p2 -= 1
      b[len(b) - 1 - i] = v2
      if p2 < 0 {
        finish = p1
        fi = i
        break
      }
    }
  }

  for i := finish; i > 0; i-- {
    fi += 1
    if fi >= len(b) {
      break
    }
    b[len(b) - 1 - fi] = s[i]
  }

  copy(s, b)
}
