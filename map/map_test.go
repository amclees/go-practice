package hmap

import "testing"

type intKey int

func (key intKey) Hash() int {
  return int(key)
}

func TestMap(t *testing.T) {
  m := Map{}
  m.Init(4)

  k1 := Key(intKey(25))
  a := interface{}(1)
  v1 := &a
  m.Put(&k1, v1)

  _, val := m.Get(&k1)
  if *val != a {
    t.Errorf("Expected to get %d from map at k1, got %d", a, *val)
  }
}
