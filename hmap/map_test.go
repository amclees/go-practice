package hmap

import (
	"math/rand"
	"testing"
)

type intKey int

func (key intKey) Hash() int {
	return int(key)
}

func TestMapStorage(t *testing.T) {
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

func TestMapResize(t *testing.T) {
	m := Map{}
	m.Init(100)

	s := make([]Value, 1000)

	for i := range s {
		k := Key(intKey(i))
		a := interface{}(rand.Int())
		v1 := &a
		m.Put(&k, v1)
		s[i] = Value{Key: &k, Val: v1}
	}

	for i, v := range s {
		k := Key(intKey(i))
		found, val := m.Get(&k)
		if !found {
			t.Errorf("Expected to find value at key %d", (*v.Key).Hash())
			continue
		}
		if *val != *v.Val {
			t.Errorf("Expected to get %d from map at key %d, got %d", *v.Val, (*v.Key).Hash(), *val)
		}
	}
}

func TestMapRemove(t *testing.T) {
	m := Map{}
	m.Init(4)

	k1 := Key(intKey(25))
	a := interface{}(1)
	v1 := &a
	m.Put(&k1, v1)

	found, _ := m.Get(&k1)
	if !found {
		t.Errorf("Expected to find value before deletion")
	}

	removed := m.Remove(&k1)
	if !removed {
		t.Errorf("Expected removal to succeed")
	}

	foundAfterRemove, _ := m.Get(&k1)
	if foundAfterRemove {
		t.Errorf("Expected not to find value after deletion")
	}
}
