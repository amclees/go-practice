package hmap

type Key interface {
  Hash() int
}

type Value struct {
  Val *interface{}
  Key *Key
}

type Map struct {
  s [][]Value
}

func (m *Map) Init(cap int) {
  m.s = make([][]Value, cap)
  for i := range m.s {
    m.s[i] = make([]Value, 1)
  }
}

func (m *Map) Put(k *Key, d *interface{}) bool {
  v := Value{Val: d, Key: k}
  ssi := m.compress((*k).Hash())
  ss := m.s[ssi]

  ni := -1
  for i, sv := range ss {
    if sv == v {
      return false
    } else if sv.Val == nil {
      ni = i
    }
  }

  if ni != -1 {
    ss[ni] = v
    return true
  }

  nss := make([]Value, newSize(ss))
  copy(nss, ss)
  nss[len(ss)] = v
  m.s[ssi] = nss
  return true
}

func (m *Map) Get(k *Key) (bool, *interface{}) {
  for _, sv := range m.s[m.compress((*k).Hash())] {
    if sv.Key == nil {
      continue
    }
    if (*sv.Key).Hash() == (*k).Hash() {
      return true, sv.Val
    }
  }
  return false, nil
}

func (m *Map) compress(hashcode int) int {
  return hashcode % len(m.s)
}

func newSize(s []Value) int {
  return len(s) + 1
}
