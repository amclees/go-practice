package stack

type Stack struct {
  s []int64
}

func (st *Stack) Init(cap int) {
  st.s = make([]int64, 0, cap)
}

func (st *Stack) Push(d int64) {
  st.s = append(st.s, d)
}

func (st *Stack) PushManual(d int64) {
  if cap(st.s) - len(st.s) >= 1 {
    st.s = st.s[:len(st.s) + 1]
  } else {
    s := make([]int64, len(st.s) + 1)
    copy(s, st.s)
    st.s = s
  }
  st.s[len(st.s) - 1] = d
}

func (st *Stack) Pop() int64 {
  last := len(st.s) - 1
  d := st.s[last]
  st.s = st.s[:last]
  return d
}

func (st *Stack) Grow(cap int) {
  s := make([]int64, cap + len(st.s))
  copy(s, st.s)
  st.s = s
}
